#!/bin/bash

while getopts n: option 
do 
    case "${option}" 
    in 
    n) NAME=${OPTARG};; 
    esac 
done 

echo "hello $NAME"