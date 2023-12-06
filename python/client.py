from proto import api_pb2
from proto import api_pb2_grpc
import grpc
import time


def client_function():

    with grpc.insecure_channel('127.0.0.1:5002') as channel:
        stub = api_pb2_grpc.AdditionStub(channel)
        print("Create a call to the server : ")
        print("1. Send two numbers and get sum.")
        print("2. Send two numbers and get sum but 3 times in a stream.")

        userInput = input("Select from above : ")
        
        if userInput =="1":
            firstNumber = int(input("Enter first number : "))
            secondNumber = int(input("Enter second number : "))
            sendRequest = api_pb2.request(a = firstNumber, b = secondNumber)
            print("Sent data and recieved response : ", stub.add(sendRequest))
        
        
        if userInput =="2":
            
            firstNumber = int(input("Enter first number : "))
            secondNumber = int(input("Enter second number : "))
            sendRequest = api_pb2.request(a = firstNumber, b = secondNumber)

            for response in stub.addStreamer(sendRequest):
                print("I got a response : ", response)
                time.sleep(2)
                



client_function()