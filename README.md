# inputFileProject
Create a project that calls service created above, pass text from input file ‘GoLang_Test.txt’ and prints JSONoutput returned from the service. 

# EndPoint
POST API: localhost:8099/upload

# Request
upload a .txt file from postman as a form-data with key GoLang_Test

# Response
"[{"word":"in","count":2774},{"word":"ut","count":1846},{"word":"dolor","count":1846},{"word":"dolore","count":1844},{"word":"et","count":931},{"word":"est","count":925},{"word":"id","count":924},{"word":"sint","count":924},{"word":"non","count":924},{"word":"qui","count":924}]"
