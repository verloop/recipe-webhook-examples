
def card(*args, title, subtitle, image_url=None, buttons=None):
    resp = {
        "title": title,
        "subtitle": subtitle,
    }
    if image_url:
        resp["image_url"] = image_url

    if buttons:
        resp["buttons"] = buttons
    return resp


def action(*args, next_block="", **variables):
    act = None
    if variables or next_block:

        act = {
            "next_block": next_block,
            "variables": {k: str(v) for k, v in variables.items()}
        }
    return act


def quick_reply(*args, title, action=None, icon=None):
    resp = {
        "title": title,
    }
    if action:
        resp["action"] = action

    if icon:
        resp["icon"] = icon
    return resp


def postback(*args, title, action=None):

    resp = {
        "type": "postback",
        "title": title
    }

    if action:
        resp["action"] = action
    return resp


def url(*args, title, url):
    return {
        "type": "web_url",
        "title": title,
        "url": url
    }
