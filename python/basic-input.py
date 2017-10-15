def reply(userTxt):
    if type(userTxt) == type('str'):
        return "The User said : "+userTxt
    else:
        return None

userTxt = raw_input('Please Say Something: ')
machineReply = reply(userTxt)

print(machineReply)
