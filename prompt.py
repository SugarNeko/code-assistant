class Prompt(object):
    def __init__(self):
        self.messages = [
            {
                "role": "system",
                "content": "You are a helpful code assistant that can teach a junior developer how to code. Don't explain the code, just generate the code block itself."
            },
            {
                "role": "user",
                "content": "",
            }
        ]


    def test_config(self, connect_detail, file_name,
                    rpc, basic_info, request_message, response_message,
                    job_details,
                    ):

        self.messages[1]['content'] = f'''Your task is to generate Go test code to test the gRPC service using the testing package.
Just generate the code block itself!!!

Here are the requirements:
- Use standard testing package
- Include client response validation
- Test server response validation
- Connect Timeout should be 15 seconds

{job_details}

Connect detail: {connect_detail}
Proto file in proto/{file_name} and code generate in the same directory as proto file
Sample proto reference:

{basic_info}

{rpc}

{request_message}

{response_message}
        '''
        # print(self.messages[1]['content'])
        return self.messages

    def job_details(
            self,
            full_parameters_check, specific_parameters_check, specific_field_check, error_message_check,
            request_parameters, response_field
    ):

        text = "Verification Details:\n\n"
        if full_parameters_check:
            text += "- Positive testing: Constructing typical requests that fully comply with the interface specification to verify whether the complete request parameters is executed correctly\n"
        if specific_parameters_check:
            no = 1
            text += "- Validate given special request parameters\n"
            for key, value in request_parameters.items():
                text += f"  {no}. {key}={value}\n"
                no += 1
        if specific_field_check:
            no = 1
            text += "Verify that the given response field is correct\n"
            for key, value in response_field.items():
                text += f"  {no}. {key}={value}\n"
                no += 1

        return text.rstrip("\n")


