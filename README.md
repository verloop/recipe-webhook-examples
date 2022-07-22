This repository contains examples of writing a HTTP API for
using webhook blocks in a Verloop recipe.

[`python`](python/) directory has an example of setting up a HTTP server in python. The
server is based on Flask and you have to host it on your own.

[`js`](js/) directory has an example of setting up a HTTP server in Javascript.
The example uses Firebase cloud functions to host the serverless code.

[`golang`](go/) directory has an example of setting up a HTTP server in Go lang.
The example uses `net/http` package.

## Request Format
Verloop server sends the request to your HTTP server in the following JSON format.
```javascript
{
  "variables": {
    "operator": {
      "type": "TEXT",
      "value": "Airtel",
      "parsed_value": "Airtel"
    },  
    "rechargeNumber": {
      "type": "TEXT",
      "value": "9988776655",
      "parsed_value": "+919988776655"
    },
  },
  "current_block": "Get_Plans",
  "visitor": {
    "name": "Vinod",
    "email": "",
    "phone": "",
    "avatar": ""
  },
  "state": {},
  "visitor_custom_fields": {
    "customerId": "ABDBA905636"
  }, 
  "room_custom_fields": {
    "key": "value"
  }
}
```

* `variables`: A JSON object having all the variables defined in a recipe. The `type`,
`value` and `parsed_value` of the variable is also sent in the request.
In the example above, `operator` and `rechargeNumber` are the two variables defined the
recipe.

* `current_block`: Name of the current webhook block being executed.

* `visitor`: A JSON object having all fields related to the visitor of your website.

* `state`: A JSON object which can hold custom key, value pairs. The value of this JSON object has to be set in the response. The same state will be returned in the
next webhook request. The state is always empty for the first webhook request of
the conversation.

* `visitor_custom_fields`: A JSON object which can hold custom key, value pairs. You can set the custom fields for the visitor by customizing the Verloop widget installation script.

* `room_custom_fields`: A JSON object which can hold custom key, value pairs. You can set the custom fields for the room/conversation by customizing the Verloop widget installation script.


## Response Format
Verloop expects the response to be in the following format.
```javascript
{
  "next_block":"Welcome",
  "variables": {
    "operator": "Airtel",
    "rechargeNumber": "+919988776655",
    "sys_language" : ""
  },
  "state": {
    "order_id": "GuyrHft6FHyeur72"
  },
  "exports": {
    "OperatorList": [
      {
        "title": "Airtel",
        "action": {
          "next_block": "Show_Plans",
          "variables": {
            "operator": "Airtel"
          }
        },
        "triggers":["The 4G Girl"]
      }
      {
        "title": "Vodafone",
        "action": {
          "next_block": "Show_Plans",
          "variables": {
            "operator": "Vodafone"
          }
        },
        "triggers":["Zoozoo","Pug"]
      }
    ]
  }
}
```

* `next_block`: The name of the block to be executed next. This will override the
next block set by the Recipe builder interface. If `next_block` is not present
in the response, bot goes to next block set by the Recipe builder interface.

* `variables`: A JSON object having key value pairs. The values of the variables will
be updated to reflect these values. All subsequent blocks will see the updated
values of the variables.
In the above example, two variables are being set: `operator` and `rechargeNumber`
Note: All the variables should be first declared in the Recipe builder interface.
  * `sys_language` is a system variable of type `Language`. It holds the current language in which the chat is being run. Webhooks can change the value of `sys_language` to available language name or id in the recipe and the subsequent chat flow will change to that language.

* `state`: A JSON object which can hold custom key, value pairs. The same state is
returned in the subsequent webhook requests. In the sample response above, an `orderid`
key is being set. These key values cannot be referenced through the Recipe builder UI.

* `exports`: A JSON object having the details of all the templates to be created.
In the above response, a template `OperatorList` is being created.
Its a `Quick Reply` template. See list of all possible [types of templates](#types-of-templates).
These templates can be used in the subsequent blocks. The templates and their types
have to declared while configuring the webhook block in the Recipe builder interface. If no template have to be returned any empty dictionary should be retuned.

## Types of Templates
### Quick Reply Template
Quick reply template has to be in the following format. It should contain list
of all the quick reply options.
```javascript
[
  {
    "title": "Title",
    "action": {
      "next_block": "Block_Name",
      "variables": {
        "variable_name": "variable_value",
        "another_variable_name": "another_variable_value"
      }
    },
    "triggers":["trigger", "list"]
  },
  {
    "title": "Title",
    "action": {
      "next_block": "Block_Name",
      "variables": {
        "variable_name": "variable_value",
        "another_variable_name": "another_variable_value"
      }
    },
    "triggers":["trigger", "list"]
  }
]
```
  * `title`: The title of the quick reply option to be shown
  * `action`: An action to take when the user clicks on this quick reply option.
    * `next_block`: Which block to go to when the user clicks this option
    * `variables`: The name and value of variables to be set. Multiple variables
    can be set here.
  * `triggers`: List of words which triggers the next_block when the matching words is entered by user.

### Buttons Template
This template can be used to construct a Buttons block.
The format should be as shown below
```javascript
[
  {
      "type":"postback",
      "title":"Buy",
      "action":{
         "next_block":"Do_Payment",
         "variables":{
            "paymentInitiated": "Yes"
         }
      }
   },
   {
      "type":"web_url",
      "title":"Visit Website",
      "url":"https://verloop.io"
   }
]
```
  * `type`: The type of the button. It can be `postback` or `web_url`.
    * `postback`: This indicates that when the user clicks this button, take him
    to another block. The next block to go to and the variables to be set are
    specified in the `action` object.
    * `web_url`: When the user clicks this button, another window opens with
    the specified url.
  * `title`: Title of the button to be shown.
  * `action`: Action to take when the user clicks this button. Valid only for
  buttons of type `postback`
  * `url`: A URL to go to when the user clicks this button.
  Valid only for buttons of type `web_url`

### Slider Template (Also called as Cards Template)
This template can be used in constructing a Slider block. The below example
initialises a slider with a single slide in it.
```javascript
[
  {
    "title":"Product Name",
    "subtitle":"product details",
    "image_url":"<url>",
    "buttons":[
        {
          "type":"postback",
          "title":"Buy",
          "action":{
             "next_block":"Do_Payment",
             "variables":{
                "paymentInitiated": "Yes"
             }
          }
       },
       {
          "type":"web_url",
          "title":"Know More",
          "url":"https://verloop.io"
       }
    ]
  }
]
```
  * `title`: Heading of this card
  * `subtitle`: Sub heading
  * `image_url`: The url of the image to be shown
  * `buttons`: List of buttons for this card. Maximum of three buttons are allowed
  for a card. Each button can of type `postback` or `web_url`. See the [Buttons Template](#types-of-templates)
  section to understand how to create buttons.


### List Template
This template can be used in constructing a List block. The below example
initialises a list with multiple sections with each section containing multiple items in it.

```javascript
[{
    "title": "Airtel Voice",
    "items": [{
        "title": "Voice Pack 1",
        "subtitle": "subtitle 1"
      },
      {
        "title": "Voice - Diwali Offer",
        "subtitle": "subtitle 2"
      }
    ]
  },
  {
    "title": "Airtel Data",
    "items": [{
        "title": "Data Pack - 30GB",
        "subtitle": "sub1"
      },
      {
        "title": "Data - Unlimited",
        "subtitle": "sub 2"
      }
    ]
  }
]
```

Note: Character length of `title` and `subtitle` cannot exceed 24 and 72 respectively.