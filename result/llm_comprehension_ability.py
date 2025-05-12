from azure.ai.inference import ChatCompletionsClient
from azure.ai.inference.models import SystemMessage, UserMessage
from azure.core.credentials import AzureKeyCredential

import yaml

with open("config.yaml", "r", encoding="utf-8") as file:
    config = yaml.safe_load(file)

token = config["provider"]["github"]
endpoint = "https://models.github.ai/inference"
models = [
    "openai/gpt-4.1", "openai/gpt-4o",
    "meta/Meta-Llama-3.1-8B-Instruct", "meta/Meta-Llama-3.1-70B-Instruct",
    # "meta/Meta-Llama-3.1-405B-Instruct",
    # "meta/Llama-4-Scout-17B-16E-Instruct"
    "deepseek/DeepSeek-V3-0324",
    "deepseek/DeepSeek-R1"
]

for model in models:

    client = ChatCompletionsClient(
        endpoint=endpoint,
        credential=AzureKeyCredential(token),
        token=4096,
    )

    messages = [
        SystemMessage("You are a helpful code assistant that can teach a junior developer how to code."),
        UserMessage(
            '''
    Analyze the following Protobuf and answer the following questions:
    
    1. How should the Go code import the code generated from this Protobuf?
    2. How should the client-side Go code call the GetUser RPC service?
    3. What field constraints have been applied to the UserRequest in the request?
    
    syntax = "proto3";
    
    package hello.v2;
    
    import "google/protobuf/timestamp.proto";
    import "github.com/envoyproxy/protoc-gen-validate/validate/validate.proto";
    
    option go_package = "code-assistant/proto/example";
    
    service UserService {
        rpc GetUser (UserRequest) returns (UserResponse);
        ...... // 
    }
    
    message UserRequest {
        int32 user_id = 1  [(validate.rules) = {int32: {gte: 1}}];
    }
    
    message UserResponse {
        string name = 1;
        string email = 2;
    }
            '''
        ),
    ]

    response = client.complete(messages=messages, model=model)
    content = response.choices[0].message.content
    if model == "deepseek/DeepSeek-R1":
        content = content.split("</think>")[-1]
    print(content)

    with open(f'result/llm_test_1/{model.split("/")[1]}.md', 'w') as f:
        f.write(content)


