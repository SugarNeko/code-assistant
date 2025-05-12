from proto_schema_parser.parser import Parser, ast
from proto_schema_parser.generator import Generator
from typing import Any


class Interpreter(object):
    def __init__(self):
        self.file = ast.File

    def load_file(self, content):
        self.file = Parser().parse(content)
        # for element in self.file.file_elements:
        #     print(element)
        return self

    def full_file_name(self, name: str):
        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Package':
                file_name = element.name
                file_name = file_name.replace('.', '/')
                return f'{file_name}/{name}'

    def package(self):
        elements = []

        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Package':
                elements.append(element)

        return Generator().generate(ast.File(file_elements=elements))

    def imports(self):
        imports = []
        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Import':
                imports.append(element.name)

        return imports

    def import_content(self):
        elements = []

        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Import':
                elements.append(element)

        return Generator().generate(ast.File(file_elements=elements))

    def options(self):
        elements = []

        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Option':
                elements.append(element)

        return Generator().generate(ast.File(file_elements=elements))

    def message(self, message_name: str):
        elements = []
        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Message' and message_name == element.name:
                elements.append(element)

        return Generator().generate(ast.File(file_elements=elements))

    def message_elements(self, message_name: str):
        for e1 in self.file.file_elements:
            if e1.__class__.__name__ == 'Message' and message_name == e1.name:
                return e1.elements

        return None

    def services(self):
        services = []

        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Service':
                # print(element.name)
                for node in element.elements:
                    if node.__class__.__name__ == 'Method':
                        services.append({
                            'service': element.name,
                            'rpc': node,
                        })

        return services

    def rpc(self, service_name: str, rpc_name: str):
        elements = []

        for e1 in self.file.file_elements:
            if e1.__class__.__name__ == 'Service' and e1.name == service_name:
                service = ast.Service(name=e1.name)
                elements.append(service)
                for e2 in e1.elements:
                    if e2.__class__.__name__ == 'Method':
                        if e2.name == rpc_name:
                            service.elements.append(e2)
                            # print(e2)
                            return Generator().generate(ast.File(file_elements=elements)), e2.input_type.type, e2.output_type.type
                        else: service.elements = []
                    else:
                        service.elements.append(e2)


    def basic_info(self):
        basic = ast.File(
            syntax="proto3",
        )

        for element in self.file.file_elements:
            if element.__class__.__name__ in ['Package', 'Import', 'Option', 'Comment']:
                basic.file_elements.append(element)
            else:
                break

        return Generator().generate(basic)

    def go_package(self):
        elements = []
        for element in self.file.file_elements:
            if element.__class__.__name__ == 'Package' and element.name == 'go_package':
                elements.append(element)