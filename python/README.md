A Flask based web server for speedy development of end points to serve Verloop recipe.
To know about the formats of request and response see [this](../README.md)

# How to use?

## Configure server
In the Verloop recipe builder, add a webhook block with name `name_of_block` on the recipe and set `https://<YOU_SERVER_URL>/my_bot` as the webhook URL.

Add a file `my_bot.py` in the `bots` folder with the endpoint you want with a `Bot` class. The class name *must* be `Bot`. A method `name_of_block` should be present in the `Bot` class.

## Examples
Checkout an example at [`bot/rechargebot.py`](./bots/rechargebot.py)
This example is to be used with the Verloop recipe template `Recharge Bot (Using Webhooks)`
