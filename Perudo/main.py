

def make_guess(my_dices):
    return []

def valid_turn(legend, new_guess, maputo=False):
    if new_guess[1] > 6:
        print('dice cant be more then 6')
        return False
    
    last_guess = legend[-1]

    if maputo:
        if new_guess[0] == last_guess[0]:
            
        elif new_guess[0] > last_guess[0]:

    else:
        if last_guess[1] == 1 and new_guess[1] == 1:
            if new_guess[0] > last_guess[0]:
                return True

        elif last_guess[1] == 1 and new_guess[1] != 1:
            if new_guess[0] >= last_guess[0] * 2 + 1:
                return True
        
        elif last_guess[1] != 1 and new_guess[1] == 1:
            if new_guess[0] >= (last_guess[0] // 2 + last_guess[0] % 2):
                return True

        elif last_guess[1] != 1 and new_guess[1] != 1:
            if new_guess[0] > last_guess[0]:
                return True
            
            if new_guess[0] == last_guess[0] and new_guess[1] > last_guess[1]:
                return True
        
        return False

                

legend = []
new_guess = list(map(int, input('new guess: ').split()))
legend.append(new_guess)

while True:
    new_guess = list(map(int, input('new guess: ').split()))
    valid = valid_turn(legend=legend, new_guess=new_guess, maputo=False)
    print(valid)
    if valid:
        legend.append(new_guess)






exit()

NUM_OF_PLAYERS = -1
NUM_OF_DICES = -1

print('Welcome to perudo game helper')

if input('DEFAULT?: y/n') == 'y':
    NUM_OF_PLAYERS = 4
    NUM_OF_DICES = 6
else:
    NUM_OF_PLAYERS = input('Enter num of players: ')
    NUM_OF_DICES = input('Enter num of dices: ')
    if NUM_OF_DICES <= 0 or NUM_OF_PLAYERS <= 0:
        print('Error in settings')
        exit()

while True:
    legend = []
    my_dices = []
    for i in range(NUM_OF_DICES):
        my_dices.append(input('My dice {}/{} : '.format(i+1, NUM_OF_DICES)))


    while True:
        if input('My turn?: y/n') == 'y':
            my_guess = make_guess(my_dices)
            legend.append(my_guess)
        else:
            else_guess = list(map(int, input('Else guess').split()))
            print(else_guess)
            legend.append(else_guess)

    

    

    