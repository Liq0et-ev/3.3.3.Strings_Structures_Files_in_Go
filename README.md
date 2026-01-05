# 3.3.3.Strings_Structures_Files_in_Go
# Iekšzemes autobusu maršrutu apstrāde 

Šī Go programma nolasa informāciju par iekšzemes autobusu maršrutiem no teksta faila `db.csv`, kas atrodas tekošajā mapē un ir UTF-8 kodējumā.<br>
<br>
Katrs maršruts satur laukus: sākuma pietura, gala pietura, nedēļas diena (Pr, Ot, Tr, Ce, Pt, St, Sv), laiks (hh:mm) un biļetes cena (reāls skaitlis ar 2 cipariem aiz punkta).<br>
<br>
Failā lauki ir atdalīti ar komatiem, bet ieraksti - ar jaunās rindas simboliem; iespējamas tukšas rindas un liekas atstarpes pirms/pēc lauku vērtībām, kuras programma ignorē.<br>
<br>
Programma veic datu validāciju: ja rindā trūkst lauku, lauku ir par daudz vai datu tipi/formāti nav pareizi, šī rinda netiek iekļauta rezultātā, bet tiek ierakstīta failā `err.txt` (tajā pašā mapē).<br>
<br>
Visi veiksmīgi nolasītie un apstrādātie maršruti tiek saglabāti masīvā/segmentā no struktūrām un pēc tam ierakstīti JSON failā `bus.json` tekošajā mapē.
