import json

import requests
from openai import OpenAI
from azure.ai.inference import ChatCompletionsClient
from azure.ai.inference.models import SystemMessage, UserMessage
from azure.core.credentials import AzureKeyCredential
import yaml

with open("config.yaml", "r", encoding="utf-8") as file:
    config = yaml.safe_load(file)


def openai_provider(model_name, messages):
    client = OpenAI(
        api_key=config['provider']['openai'],
    )

    response = client.responses.create(
        model=model_name,
        input=messages,
    )

    return response.output_text

def github_provider(model_name, messages):
    client = ChatCompletionsClient(
        endpoint="https://models.github.ai/inference",
        credential=AzureKeyCredential(config['provider']['github']),
    )

    response = client.complete(
        messages=messages,
        temperature=1.0,
        top_p=1.0,
        max_tokens=4096,
        model=model_name
    )

    return response.choices[0].message.content

def deepseek_provider(model_name, messages):
    client = OpenAI(
        api_key=config['provider']['deepseek'], base_url="https://api.deepseek.com"
    )

    response = client.chat.completions.create(
        model=model_name,
        messages=messages,
        stream=False
    )

    return response.choices[0].message.content


def x_ai_provider(model_name, messages):
    client = OpenAI(
        api_key=config['provider']['xai'],
        base_url="https://api.x.ai/v1",
    )

    completion = client.chat.completions.create(
        model=model_name,
        messages=messages,
    )

    return completion.choices[0].message.content

def custom_model(messages):
    r = requests.post("http://10.0.1.228:6006/", data=json.dumps({'chat': messages}))
    # print(r.text)
    return r.json()['response']