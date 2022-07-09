import time
import pymysql.cursors

conn = pymysql.connect(
   host='127.0.0.1',
   user='root',
   password='12345',
   database='ebook',
   charset='utf8mb4',
   cursorclass=pymysql.cursors.DictCursor
)

print("hello")

cursor = conn.cursor()

sql = """INSERT INTO user_tables (name,password) VALUES("stress_test_{}", "1234567890");"""

print(sql)

# tmp_sql = sql.format(99999)
# print("tmp_sql:",tmp_sql)
# cursor.execute(tmp_sql)
for i in range(10000000, 20000000):
   try:
      tmp_sql = sql.format(i)
      cursor.execute(tmp_sql)
      if not i % 10000:
         conn.commit()

   except Exception as e:
      print(e)

# Closing the connection
conn.commit()
conn.close()
