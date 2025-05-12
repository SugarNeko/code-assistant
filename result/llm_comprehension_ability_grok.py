from openai import OpenAI
import yaml

with open("config.yaml", "r", encoding="utf-8") as file:
    config = yaml.safe_load(file)

messages = [
    {
        "role": "system",
        "content": "You are a helpful code assistant that can teach a junior developer how to code."
    },
    {
        "role": "user",
        "content": '''
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
            ''',
    }
]

client = OpenAI(
    api_key=config['provider']['xai'],
    base_url="https://api.x.ai/v1",
)

completion = client.chat.completions.create(
    model='grok-3-latest',
    messages=messages,
)

content = completion.choices[0].message.content
print(content)

with open(f'result/llm_test_1/Grok-3.md', 'w') as f:
    f.write(content)
