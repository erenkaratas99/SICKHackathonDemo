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

type_inpute_msg = "You can type your request in here (type 'quit' to terminate): "

not_und_msg = "I do not understand, can you try again?\n "

yn_response = "If the data you want is below type 'yes', otherwise type anything.\n"

purchase_id_msg = "Hello, please type the last two numbers of 'purchase order ID' to process: "



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
    customer_ID = "45000000" + input("Hello, please type the last two numbers of 'purchase order ID' to process: ")
    word_file = open("keywords.txt", "w")
    kword_file.write(purchase_id_msg)
    kword_file.close()
    kword_file = open("keywords.txt", "r")
    while True:
        if (customer_ID[-1].isnumeric() and customer_ID[-2].isnumeric()): 

            kword_file = open("keywords.txt", "w")
            kword_file.write(welcome_msg)
            kword_file.close()
            kword_file = open("keywords.txt", "r")

            
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
                        
                            kword_file = open("keywords.txt", "w")
                            kword_file.write(intent['responses'])
                            kword_file.close()
                            kword_file = open("keywords.txt", "r")

                            if yn_response == "yes":
                                

                                kword_file = open('keywords.txt', "w+")
                                kword_file.write(customer_ID)
                                kword_file.close()

                                line = "\n" + str(intent['responses'])[1:-1]
                                with open("keywords.txt", "a+") as f:
                                    f.writelines(line)
                                # os._exit(os.EX_OK)
                                break
                        
                    
                else:
                    kword_file = open("keywords.txt", "w")
                    kword_file.write(not_und_msg)
                    kword_file.close()
                    kword_file = open("keywords.txt", "r")


        else:
            print("Please try again!")
            break

