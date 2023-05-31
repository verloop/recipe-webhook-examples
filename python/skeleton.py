import sys
import os
import argparse
import jinja2

cwd = os.getcwd()
base_path = os.path.dirname(os.path.realpath(__file__))


# Jinja2 environment
def render_template(template_name, template_var, client_id=None):
    searchpath = [
        os.path.join(base_path, "templates"),
    ]
    if client_id:
        searchpath.append(
            os.path.join(base_path, f"src/local_services/{client_id}"),
        )

    template_loader = jinja2.FileSystemLoader(searchpath=searchpath)
    template_env = jinja2.Environment(loader=template_loader)

    template = template_env.get_template(template_name)
    return template.render(template_var)


# importing here to avoid circular import. Dont change.
from scaffolding import generate_base, generate_llm, generate_apis


def main(argv):
    parser = argparse.ArgumentParser(description="Scaffolding a bots structure.")
    parser.add_argument(
        "-c",
        "--client_id",
        help="The client id for which you would be generating files.",
        required=True,
    )
    parser.add_argument(
        "-rt",
        "--recipe_type",
        help="""The recipe type for your client
        "base": for base recipe type.
        "llm": for llm recipe type.
        "api": for generating api methods in service files (curl request required)
        Default base recipe will be generated if no argument is specified.
        No recipe will be generated on an invalid arguemnt.
        """,
    )

    parser.add_argument(
        "-cf",
        "--curl_file",
        type=argparse.FileType("r"),
        help="""path to file with curl requests (separated by blank lines)""",
        required=False,
    )

    args = parser.parse_args()
    client_id = args.client_id
    recipe_type = args.recipe_type
    curl_file = args.curl_file

    template_var = {
        "base_client_id": client_id,
        "client_id": client_id[0].upper() + client_id[1:],
    }

    # generate files and folders for specific recipe_types
    if recipe_type == "base":
        generate_base.base(template_var=template_var)
        print("Base recipe generated.")

    elif recipe_type == "llm":
        generate_llm.llm(template_var=template_var)
        print("llm recipe generated.")

    elif recipe_type == "api":
        if not curl_file:
            print("curl file required!")
        else:
            generate_apis.api(
                template_var=template_var, curl_texts=curl_file.read().split("\n\n")
            )
            print("api services generated.")

    elif recipe_type is None:
        generate_base.base(template_var=template_var)
        print("Default recipe generated.")

    else:
        sys.exit("Invalid recipe_type specified. No files generated.")


if __name__ == "__main__":
    main(sys.argv)
