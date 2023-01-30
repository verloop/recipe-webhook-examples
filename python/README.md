A Flask based web server for speedy development of end points to serve Verloop recipe.
To know about the formats of request and response see [this](../README.md)

# How to use?

## Configure server

In the Verloop recipe builder, add a webhook block with name `name_of_block` on the recipe and set `https://<YOU_SERVER_URL>/my_bot` as the webhook URL.

Add a file `my_bot.py` in the `bots` folder with the endpoint you want with a `Bot` class. The class name _must_ be `Bot`. A method `name_of_block` should be present in the `Bot` class.

## Examples

Checkout an example at [`bot/rechargebot.py`](./bots/rechargebot.py)
This example is to be used with the Verloop recipe template `Recharge Bot (Using Webhooks)`

## Scaffolding(kinda) bots structure:

    Note: Assuming basic bots structure is like this: https://github.com/verloop/bots. Run command at the root of your directory after you clone the bots repo.

    a. `python3 skeleton.py -c credence` where credence is the client_id.

    b. Generates relevant files and some boiler plate code.

    c. We can choose to generate boilerplate code depending on the base recipe type like `LeadGenerator`, `SupportBot`, etc as available on the recipes dashboard but tbh those recipes contain static blocks and webhook blocks/functions can vary depending on the usecase(Suggestions welcomed on how to handle this.)
