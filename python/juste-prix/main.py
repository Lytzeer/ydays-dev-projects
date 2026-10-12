import random

def generate_random_number(min_value, max_value):
    return random.randint(min_value, max_value)

def main():
    max_try_number=10
    try_number = 0
    min_value = 1
    max_value = 100
    random_number = generate_random_number(min_value, max_value)

    print(f"Devine le nombre entre {min_value} et {max_value}")
    while try_number < max_try_number :
        user_number = input("> ")
        if not user_number.isdigit():
            print("Veuillez entrer un nombre valide.")
            continue
        user_number = int(user_number)

        if user_number < min_value or user_number > max_value:
            print(f"Veuillez entrer un nombre entre {min_value} et {max_value}.")
            continue

        try_number+=1

        if user_number == random_number:
            print(f"Bravo, trouvé en {try_number} coups !")
            return
        elif user_number > random_number:
            print("C'est moins !")
        else:
            print("C'est plus !")



if __name__ == "__main__":
    main()