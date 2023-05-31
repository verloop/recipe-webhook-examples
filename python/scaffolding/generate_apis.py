import os
import sys
import re

import uncurl

from skeleton import base_path, render_template


def api(template_var: dict, curl_texts: list):
    try:
        reqs = []
        for curltext in curl_texts:
            curltext = re.sub(r"\\\n", "", curltext)
            curltext = (
                curltext.replace(" --location", "")
                .replace("--form", "--data")
                .replace("file=@", "")
            )

            req = uncurl.parse_context(curltext)
            method_name = req.url.split("/")[-1]
            method_name = method_name.split("?")[0]

            endpoint = method_name.split("?")[0]
            params = {}
            try:
                params_str = method_name.split("?")[1]
                params_obj = params_str.split("&")
                for param in params_obj:
                    param = param.split("=")
                    params[param[0]] = param[1]
            except Exception as e:
                sys.exit(e)

            if not any(ch.isupper() for ch in method_name):
                method_name = req.method.lower() + "_" + method_name

            headers = dict(req.headers)
            if "application/json" in headers.values():
                isjson = True
                headers.pop(
                    list(headers.keys())[
                        list(headers.values()).index("application/json")
                    ],
                    None,
                )
            else:
                isjson = False

            if isjson:
                data_field = "data"
            else:
                data_field = "json_data"

            reqs.append(
                dict(
                    method_name=method_name,
                    url=req.url,
                    endpoint=endpoint,
                    params=params,
                    headers=req.headers,
                    method=req.method,
                    data=req.data,
                    data_field=data_field,
                )
            )

        path = base_path + "/src/local_services/" + template_var.get("base_client_id")
        try:
            os.makedirs(path, 0o755)
        except:
            pass
        file_exists = os.path.isfile(path + "/service.py")
        file_content = None
        if file_exists:
            with open(os.path.join(path, "service.py"), "r") as file:
                file_content = file.read()

        # generate service.py
        with open(os.path.join(path, "service.py"), "w") as fd:
            template_var.update(
                requests=reqs,
                file_exists=file_exists,
                existing_file_content=file_content,
            )
            fd.write(
                render_template(
                    "service_with_apis.jinja2",
                    template_var,
                    template_var.get("base_client_id"),
                )
            )

    except Exception as e:
        sys.exit(e)
