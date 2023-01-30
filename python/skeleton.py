import sys
import os
import argparse
import jinja2

cwd = os.getcwd()
base_path = os.path.dirname(os.path.realpath(__file__))

# Jinja2 environment
template_loader = jinja2.FileSystemLoader(
    searchpath=os.path.join(base_path, "templates")
)
template_env = jinja2.Environment(loader=template_loader)


def render_template(template_name, template_var):
    template = template_env.get_template(template_name)
    return template.render(template_var)


def main(argv):

    parser = argparse.ArgumentParser(description="Scaffolding a bots structure.")
    parser.add_argument(
        "-c",
        "--client_id",
        help="The client id for which you would be generating files.",
    )
    args = parser.parse_args()
    client_id = args.client_id

    # generate files and folders

    template_var = {
        "base_client_id": client_id,
        "client_id": client_id[0].upper() + client_id[1:],
    }

    # generate <client_id>.py
    with open(os.path.join(base_path, "src/bots/" + client_id + ".py"), "w") as fd:
        fd.write(render_template("base.jinja2", template_var))

    # make directory
    path = base_path + "/src/local_services/" + client_id
    os.makedirs(path, 0o755)

    # generate bots.py
    with open(os.path.join(path, "bots.py"), "w") as fd:
        fd.write(render_template("bots.jinja2", template_var))

    # generate service.py
    with open(os.path.join(path, "service.py"), "w") as fd:
        fd.write(render_template("service.jinja2", template_var))

    # generate renderer.py
    with open(os.path.join(path, "renderer.py"), "w") as fd:
        fd.write(render_template("renderer.jinja2", template_var))


if __name__ == "__main__":
    main(sys.argv)
