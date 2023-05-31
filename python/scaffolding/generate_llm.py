import os, sys

from skeleton import base_path, render_template


def llm(template_var: dict):
    try:
        # generate <client_id>llm.py
        with open(
            os.path.join(
                base_path, "src/bots/" + template_var.get("base_client_id") + "llm.py"
            ),
            "w",
        ) as fd:
            fd.write(render_template("llm_templates/llm.jinja2", template_var))

    except Exception as e:
        sys.exit(e)
