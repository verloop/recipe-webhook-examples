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

Note: Assuming basic bots structure is like this: https://github.com/verloop/bots. 
Run command at the root of your directory after you clone the bots repo.

a. `python3 skeleton.py -c credence -rt llm` <br>

    ❯ python3 skeleton.py --help
    usage: skeleton.py [-h] -c CLIENT_ID [-rt RECIPE_TYPE] [-cf CURL_FILE]

    Scaffolding a bots structure.

    require arguments:
    -c CLIENT_ID, --client_id CLIENT_ID
                        The client id for which you would be generating files.

    optional arguments:
    -h, --help          show this help message and exit
    -rt RECIPE_TYPE, --recipe_type RECIPE_TYPE
                        The recipe type for your client:
                            base: for base recipe type. 
                            llm: for llm recipe type.
                            api: for generating api methods in service files (curl request required) 
                        Default base recipe will be generated if no argument is specified. 
                        No recipe will be generated on an invalid arguemnt.
    -cf CURL_FILE, --curl_file CURL_FILE
                        path to file with curl requests (separated by blank lines)

b. Generates relevant files and some boiler plate code.

c. We can choose to generate boilerplate code depending on the base recipe type like `LeadGenerator`, `SupportBot`, etc as available on the recipes dashboard but tbh those recipes contain static blocks and webhook blocks/functions can vary depending on the usecase(Suggestions welcomed on how to handle this.)
