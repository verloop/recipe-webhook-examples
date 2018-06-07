const functions = require('firebase-functions');

// // Create and Deploy Your First Cloud Functions
// // https://firebase.google.com/docs/functions/write-firebase-functions
//


function get_operators() {
  return ['Airtel', 'Vodafone', 'Jio'];
}


function get_supported_operators() {
  operators_quick_replies = [];

  operators = get_operators();

  for(i = 0; i < operators.length; i++) {
    operators_quick_replies.push({
      'title': operators[i],
      'action': {
        'next_block': 'Get_Plans_Wait',
        'variables': {
          'operator': operators[i]
        }
      }
    });
  }

  return {
    'next_block': 'Show_Operators',
    'variables' : {},
    'exports': {
      'OperatorList': operators_quick_replies
    }
  }
}


function get_recharge_options(operator) {
  all_operators = get_operators()

  if (all_operators.indexOf(operator) === -1) {
    next_block = 'Invalid_Operator';
    plans = []
  }
  else {
    next_block = 'Show_Plans_Text';
    switch (operator) {
      case 'Airtel':
        plans = [
          {
            'title' : 'Data 1 GB (28 Days)',
            'subtitle' : 'Rs. 100',
            'buttons' : [
              {
                'type': 'postback',
                'title': 'Select',
                'action': {
                  'next_block': 'Do_Recharge',
                  'variables': {
                    'amount': '100'
                  }
                }
              },
              {
                'type': 'web_url',
                'title': 'Know More',
                'url': 'https://www.google.com'
              }
            ]
          },
          {
            'title' : 'Data 2 GB (28 Days)',
            'subtitle' : 'Rs. 150',
            'buttons' : [
              {
                'type': 'postback',
                'title': 'Select',
                'action': {
                  'next_block': 'Do_Recharge',
                  'variables': {
                    'amount': '150'
                  }
                }
              },
              {
                'type': 'web_url',
                'title': 'Know More',
                'url': 'https://www.google.com'
              }
            ]
          }
        ]
        break;
      case 'Vodafone':
        plans = [
          {
            'title' : 'Data 1 GB (28 Days)',
            'subtitle' : 'Rs. 100',
            'buttons' : [
              {
                'type': 'postback',
                'title': 'Select',
                'action': {
                  'next_block': 'Do_Recharge',
                  'variables': {
                    'amount': '100'
                  }
                }
              },
              {
                'type': 'web_url',
                'title': 'Know More',
                'url': 'https://www.google.com'
              }
            ]
          },
          {
            'title' : 'Data 2 GB (28 Days)',
            'subtitle' : 'Rs. 150',
            'buttons' : [
              {
                'type': 'postback',
                'title': 'Select',
                'action': {
                  'next_block': 'Do_Recharge',
                  'variables': {
                    'amount': '150'
                  }
                }
              },
              {
                'type': 'web_url',
                'title': 'Know More',
                'url': 'https://www.google.com'
              }
            ]
          }
        ];
        break;
      default:
        next_block = '';
        plans = [];
        break;
    }
  }

  return {
    'next_block': next_block,
    'variables' : {},
    'exports': {
      'PlanList': plans
    }
  };
}


function validate_options(operator, amount) {
  console.log(operator, amount)
  operators = get_operators()
  if(operators.indexOf(operator) === -1) {
    return false;
  }
  // TODO: check if amount is valid for the selected operator
  return true;
}


function get_payment_link(operator, phonenumber, amount) {
  if (validate_options(operator, phonenumber, amount)) {
    // TODO: Generate a unique payment link and order_id
    payment_link = "https://www.google.com"
    order_id = "skahjdfwousjklfn"
    return {
      'next_block': '',
      'variables': {},
      'exports': {
        'PaymentOptions': [
          {
            'type': 'web_url',
            'title': 'Pay ' + amount + ' Now',
            'url': payment_link
          },
          {
            'type': 'postback',
            'title': 'Payment Done',
            'action': {
              'next_block': 'Verify_Payment'
            }
          }
        ]
      },
      'state' : {
        'order_id': order_id
      }
    }
  } else {
    return {
      'next_block': 'Invalid_Options',
      'variables': {},
      'exports': {
        'PaymentOptions': [
          {
            'type': 'web_url',
            'title': 'Pay ' + amount + ' Now',
            'url': payment_link
          },
          {
            'type': 'postback',
            'title': 'Payment Done',
            'action': {
              'next_block': 'Verify_Payment'
            }
          }
        ]
      }
    }
  }
}


function check_payment_details(order_id) {
  // TODO: Verify that the payment has been made for order_id
  success = true;
  if (success) {
    successInfo = 'Payment is successful.'
    return {
      'next_block': 'Order_Success',
      'variables': {
        'successInfo': successInfo
      },
      'exports': {},
      'state': {}
    }
  } else {
    // TODO: Get failure reason
    failureInfo = 'Payment failed due to whatever reason'
    return {
      'next_block': 'Order_Failure',
      'variables': {
        'failureInfo' : failureInfo
      },
      'exports': {},
      'state': {}
    }
  }
}


exports.rechargebot = functions.https.onRequest((request, response) => {
  data = request.body;
  current_block = data.current_block;

  webhook_response = {}

  switch (current_block) {

    case 'Get_Supported_Operators':
      webhook_response = get_supported_operators();
      break;

    case 'Get_Plans':
      operator = undefined
      if('variables' in data) {
        if('operator' in data.variables) {
          operator = data.variables.operator.parsed_value
        }
      }

      webhook_response = get_recharge_options(operator);
      break;

    case 'Get_Payment_Link':
      operator = undefined
      amount = undefined
      phonenumber = undefined
      if('variables' in data) {
        if('operator' in data.variables) {
          operator = data.variables.operator['parsed_value']
        }

        if('amount' in data.variables) {
          amount = data.variables.amount['parsed_value']
        }

        if('rechargenumber' in data.variables) {
          phonenumber = data.variables.rechargenumber['parsed_value']
        }
      }
      console.log(operator, amount)
      webhook_response = get_payment_link(operator, phonenumber, amount)
      break;

    case 'Check_Payment':
      order_id = undefined
      if ('state' in data) {
        if ('order_id' in data.state) {
          order_id = data.state.order_id
        }
      }

      webhook_response = check_payment_details(order_id);
      break;

    default:
      break;
  }
  console.log(webhook_response);
  response.send(webhook_response);
})
