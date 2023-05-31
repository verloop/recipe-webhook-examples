import os, sys

from skeleton import base_path, render_template


def base(template_var: dict):
    try:
        # generate <client_id>.py
        with open(
            os.path.join(
                base_path, "src/bots/" + template_var.get("base_client_id") + ".py"
            ),
            "w",
        ) as fd:
            fd.write(render_template("base.jinja2", template_var))

        # make local services directory
        path = base_path + "/src/local_services/" + template_var.get("base_client_id")
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

    except Exception as e:
        sys.exit(e)
