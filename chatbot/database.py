import pandas as pd 
import csv

xml = pd.read_xml(r'D:/SICKHackathon/output.xml',xpath='//*[contains(name(),"properties")]')

xml.dropna(how='all', axis=1, inplace=True)


kywd_file = open("keywords.txt", "r")
                
variable = kywd_file.readline()

print(variable)


f = open('answer.csv', 'w')
# create the csv writer
writer = csv.writer(f)

# write a row to the csv file
writer.writerow(xml.variable)

# close the file
f.close()





