import os

from common import App, magic_router
import bots

AUTH_KEY = os.getenv("AUTH_TOKEN", "PAPA_KEHTE_HAIN_BADA_NAAM_KAREGA")


app = App(auth_key=AUTH_KEY, routes=magic_router(module=bots))


if __name__ == "__main__":
    debug = os.getenv("LOG_LEVEL", False)
    app.run(host='0.0.0.0', port=3000, debug=debug)
