import random
import names

TABLE_NAME = "User"
FIELD_NAMES = ("name", "family", "age", "sex")
SEXES = (0, 1)
MIN_AGE = 20
MAX_AGE = 80
NUM_USERS = 300

with open("insert_users.sql", "w") as f:
    first_line = 'INSERT INTO "{}" ({}) VALUES\n'.format(TABLE_NAME, ", ".join(FIELD_NAMES))
    f.write(first_line)

    value_sql = "('{name}', '{family}', {age}, {sex})"
    for i in range(NUM_USERS):
        sex = random.choice(SEXES)
        name = names.get_first_name()
        family = names.get_last_name()
        age = random.randint(MIN_AGE, MAX_AGE)
        
        values = value_sql.format(name=name, age=age, family=family, sex=sex)
        if i == NUM_USERS - 1:
            end_exec = ";"
            values += end_exec
        f.write(values + "\n")

    f.flush()
