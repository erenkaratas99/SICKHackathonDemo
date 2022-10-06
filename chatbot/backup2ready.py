from ast import Break
import random
import json
import os
from traceback import print_tb
import torch

from model import NeuralNet
from nltk_utils import bag_of_words, tokenize


welcome_msg = """

Welcome dear user, you can reach the customer datas from this dataset. 
You can check the examples before you go:
- What is the order status?
- When is the date of purchase?
- What is the supplier' information?
- I want to see payment type.
- Who's the purchaser?


"""

type_inpute_msg = "You can type your request in here (type 'quit' to terminate)\n"

not_und_msg = "I do not understand, can you try again?\n"

try_again_msg = "Please try again!\n"

order_id_msg = "Hello, please type the last two numbers of 'purchase order ID' to process.\n"

yn_response_msg = "If the data type above is correct type 'yes', otherwise type anything.\n"





device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')

with open('intents.json', 'r') as json_data:
    intents = json.load(json_data)

FILE = "data.pth"
data = torch.load(FILE)

input_size = data["input_size"]
hidden_size = data["hidden_size"]
output_size = data["output_size"]
all_words = data['all_words']
tags = data['tags']
model_state = data["model_state"]

model = NeuralNet(input_size, hidden_size, output_size).to(device)
model.load_state_dict(model_state)
model.eval()

bot_name = "CHATBOT"


while True:

    msgs_file = open("msgs.txt", "w")
    msgs_file.write(order_id_msg)
    msgs_file.close()
    msgs_file = open("msgs.txt", "r")
    print(msgs_file.read())
    customer_ID = "45000000" + input()
    welcome_status=0
    while True:
        if (customer_ID[-1].isnumeric() and customer_ID[-2].isnumeric()): 

            if welcome_status == 0:
                msgs_file = open("msgs.txt", "w")
                msgs_file.write(welcome_msg)
                msgs_file.close()
                msgs_file = open("msgs.txt", "r")
                print(msgs_file.read())
                welcome_status=1
        
            sentence = input(type_inpute_msg)
            if sentence == "quit":
                break
            else:

                sentence = tokenize(sentence)
                X = bag_of_words(sentence, all_words)
                X = X.reshape(1, X.shape[0])
                X = torch.from_numpy(X).to(device)

                output = model(X)
                _, predicted = torch.max(output, dim=1)

                tag = tags[predicted.item()]

                probs = torch.softmax(output, dim=1)
                prob = probs[0][predicted.item()]
                if prob.item() > 0.75:
                    for intent in intents['intents']:
                        if tag == intent["tag"]:


                            msgs_file = open("msgs.txt", "w")
                            msgs_file.write(yn_response_msg)
                            msgs_file.close()
                            msgs_file = open("msgs.txt", "r")
                            print(msgs_file.read())
                            

                            kywd_file = open("keywords.txt", "w")
                            kywd_file.write(str(intent['responses'])[2:-2])
                            kywd_file.close()
                            kywd_file = open("keywords.txt", "r")
                            print(kywd_file.read())

                            yn_response = input()

                            if yn_response == "yes":

                                id_file = open('id.txt', "w+")
                                id_file.write(customer_ID)
                                id_file.close()

                                kywd_file = open("keywords.txt", "w")
                                kywd_file.write(str(intent['responses'])[2:-2])
                                kywd_file.close()

                                # os._exit(os.EX_OK)
                                break
                    
                
                else:
                    msgs_file = open("msgs.txt", "w")
                    msgs_file.write(not_und_msg)
                    msgs_file.close()
                    msgs_file = open("msgs.txt", "r")


        else:
            msgs_file = open("msgs.txt", "w")
            msgs_file.write(try_again_msg)
            msgs_file.close()
            msgs_file = open("msgs.txt", "r")
            print(msgs_file.read())
            break

