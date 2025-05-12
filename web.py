import time
from datetime import datetime

import streamlit as st
import models
import proto_parser
from io import StringIO
from pathlib import Path
from prompt import Prompt


class UI(object):
    def __init__(self):
        st.session_state['proto_files'] = {}
        st.session_state['services'] = {}
        st.session_state['navigation'] = {
            "Files": [],
            "Services": [],
        }

    def start(self):
        self.show_uploader()
        return self

    def show_uploader(self):
        uploaded_files = st.sidebar.file_uploader("Choose a file", type=['proto'], accept_multiple_files=True)
        for uploaded_file in uploaded_files:
            content = StringIO(uploaded_file.getvalue().decode("utf-8")).read()
            parser = proto_parser.Interpreter().load_file(content)
            file_name = parser.full_file_name(uploaded_file.name)

            if st.session_state['proto_files'].get(file_name):
                continue

            st.session_state['proto_files'][file_name] = content
            for i in parser.imports():
                if st.session_state['proto_files'].get(i) is None:
                    st.session_state['proto_files'][i] = ''

        if len(uploaded_files) != 0:
            self.show_side_bar()

        return self

    def show_side_bar(self):
        self.update_side_bar()
        pg = st.navigation(st.session_state['navigation'])
        pg.run()

    def update_side_bar(self):
        for name, content in st.session_state['proto_files'].items():
            # Can't use '/' in url path
            url_path = name.replace('/', '_')

            # Get file content
            if content == '':
                st.session_state['navigation']['Files'].append(
                    st.Page(lambda c=content: self.proto_page(c), title='⚠️ ' + name, url_path=url_path)
                )
            else:
                st.session_state['navigation']['Files'].append(
                    st.Page(lambda c=content: self.proto_page(c), title=name, url_path=url_path)
                )

            # Get service name
            parser = proto_parser.Interpreter().load_file(content)

            for s in parser.services():
                # print(s['service'], s['rpc'].name)
                st.session_state['navigation']['Services'].append(
                    st.Page(lambda service=s: self.service_page(service), title=s['rpc'].name, url_path=s['rpc'].name)
                )
                # print(s['rpc'])
                st.session_state['services'][f"{s['service']}.{s['rpc'].name}"] = {
                    'file_name': name,
                    'service': s['service'],
                    'rpc': s['rpc'],
                }
            # pprint.pp(st.session_state['services'])
        return self

    def proto_page(self, content):
        st.markdown(f"```protobuf\n{content}\n```")

    def service_page(self, service):
        file_name = st.session_state['services'][f"{service['service']}.{service['rpc'].name}"]['file_name']
        parser = proto_parser.Interpreter().load_file(st.session_state['proto_files'][file_name])
        rpc_content, req_message_name, resp_message_name = parser.rpc(service_name=service['service'],
                                                                      rpc_name=service['rpc'].name)
        st.markdown(
            f"# {service['rpc'].name}\n"
            f"## Service\n ```protobuf\n{rpc_content}\n```\n"
            f"## Request Message\n ```protobuf\n{parser.message(req_message_name)}\n```\n"
            f"## Response Message\n ```protobuf\n{parser.message(resp_message_name)}\n```\n"
        )

        st.markdown("## Test Code Generate")

        col1, col2 = st.columns(2)

        with col1:
            model_option = st.selectbox(
                'Model selection',
                (
                    "DeepSeek-V3", "DeepSeek-R1",
                    "OpenAI GPT-4o", "OpenAI GPT-4.1",
                    "Meta-Llama-3.1-8B-Instruct", "Meta-Llama-3.1-70B-Instruct",
                    "Grok 3"
                ),

                index=0,
                placeholder='Select a model...',
            )

            connect_detail = st.text_input("Connect", "127.0.0.1:9000")

        with col2:
            st.text_input("RPC Service", service['rpc'].name)
            st.text_input("File", file_name)

        request_params, response_fields = {}, {}

        with st.expander("Request Parameters Validation"):
            full_parameters_check = st.checkbox("FULL PARAMETERS CHECK")

            specific_parameters_check = st.checkbox("SPECIFIC PARAMETERS CHECK")
            if specific_parameters_check:
                message_elements = parser.message_elements(message_name=req_message_name)
                for message_element in message_elements:
                    if message_element.__class__.__name__ != 'Field':
                        continue
                    if st.checkbox(f"{message_element.name} ({message_element.type})"):
                        message_element_value = st.text_input(
                            "Enter the value you expect 👇",
                            placeholder="Specific value",
                            key=f"param_request_{message_element.name}"  # 唯一标识符
                        )

                        request_params[message_element.name] = message_element_value
                        # print(request_params)

        with st.expander("Response Fields Validation"):
            response_fields_check = st.checkbox("RESPONSE FIELDS CHECK")
            if response_fields_check:
                message_elements = parser.message_elements(message_name=resp_message_name)
                for message_element in message_elements:
                    if message_element.__class__.__name__ != 'Field':
                        continue
                    if st.checkbox(f"{message_element.name} ({message_element.type})"):
                        message_element_value = st.text_input(
                            "Enter the value you expect 👇",
                            placeholder="Specific value",
                            key=f"param_response_{message_element.name}"  # 唯一标识符
                        )

                        response_fields[message_element.name] = message_element_value

        with st.expander("Error Messages Validation"):
            error_message_check = st.checkbox("ERROR MESSAGE CHECK")

        job_details = Prompt().job_details(
            full_parameters_check=full_parameters_check,
            specific_parameters_check=specific_parameters_check,
            specific_field_check=response_fields_check,
            error_message_check=error_message_check,
            request_parameters=request_params,
            response_field=response_fields,
        )

        with st.expander("Prompt Details"):
            prompt = Prompt().test_config(
                connect_detail=connect_detail,
                file_name=file_name,
                basic_info=parser.basic_info(),
                rpc=rpc_content,
                request_message=parser.message(message_name=req_message_name),
                response_message=parser.message(message_name=resp_message_name),
                job_details=job_details,
            )

            for p in prompt:
                if p['role'] != 'system':
                    st.markdown(f"```prorobuf\n{p['content']}\n```")

        if st.button("Generate"):
            st.success(f'Using Model: **{model_option}**')
            Path(f"result/code/{model_option.replace(" ", "-")}").mkdir(parents=True, exist_ok=True)

            for i in range(1, 31):
                with st.spinner(f'Loop {i} Waiting for it...'):
#             with st.spinner(f'Waiting for it...'):
                    if model_option == 'DeepSeek-V3':
                        response = models.deepseek_provider('deepseek-chat', prompt)
                    elif model_option == 'DeepSeek-R1':
                        response = models.deepseek_provider('deepseek-reasoner', prompt).split("</think>")[-1]
                    elif model_option == 'OpenAI GPT-4o':
                        response = models.openai_provider('gpt-4o', prompt)
                    elif model_option == 'OpenAI GPT-4.1':
                        response = models.openai_provider('gpt-4.1', prompt)
                    elif model_option == 'Meta-Llama-3.1-8B-Instruct':
                        response = models.github_provider('meta/Meta-Llama-3.1-8B-Instruct', prompt)
                    elif model_option == 'Meta-Llama-3.1-70B-Instruct':
                        response = models.github_provider('meta/Meta-Llama-3.1-70B-Instruct', prompt)
                    elif model_option == 'Grok 3':
                        response = models.x_ai_provider('grok-3-latest', prompt)
                st.markdown(response)
                with open(
                        f"result/code/{model_option.replace(" ", "-")}/{service['service']}_{service['rpc'].name}_{time.time().__int__()}_test.go",
                        'w') as c:
                    code_block = response.lstrip("```go\n").rstrip("`")
                    c.write(code_block)


UI().start()
