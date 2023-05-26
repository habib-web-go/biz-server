import random
import names
random.uniform(0, 1)
with open("insert_users.sql", "w") as f:
    line_1 = 'INSERT INTO "User" (name, family, age, sex) VALUES\n'
    f.write(line_1)
    value_sql = "('{name}', '{family}', {age}, '{sex}')"
    for i in range(100):
        sex = "male" if random.uniform(0, 1) > 0.5 else "female"
        name = names.get_first_name()
        family = names.get_last_name()
        age = random.randrange(20, 80)
        f.write(value_sql.format(name=name, age=age, family=family, sex=sex))
        if i != 99:
            f.write(",\n")
    end_exec = ";"
    f.write(end_exec)
    f.flush()

