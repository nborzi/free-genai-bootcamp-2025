## Running Ollama Third-Party Service

### Choosing a Model

You can get the model_id that ollama will launch from the [Ollama Library](https://ollama.com/library).

https://ollama.com/library/llama3.2

eg. LLM_MODEL_ID="llama3.2:1b"

### Getting the Host IP

#### Linux

Get your IP address
```sh
sudo apt install net-tools
ifconfig
```

Use this for host  `$(hostname -I | awk '{print $1}')`

HOST_IP=$(hostname -I | awk '{print $1}') NO_PROXY=localhost LLM_ENDPOINT_PORT=9000 LLM_MODEL_ID="llama3.2:1b" docker compose up


### Ollama API

Once the Ollama server is running we can make API calls to the ollama API get in github

https://github.com/ollama/ollama/blob/main/docs/api.md


## Download (Pull) a model

curl http://localhost:9000/api/pull -d '{
  "model": "llama3.2:1b"
}'

## Generate a Request

curl http://localhost:9000/api/generate -d '{
  "model": "llama3.2:1b",
  "prompt": "Why is the sky blue?"
}'

# Response 
 curl http://localhost:9000/api/generate -d '{
>   "model": "llama3.2:1b",
rompt":>   "prompt": "Why is the sky blue?"
'> }'
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:41.735339554Z","response":"The","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:41.851026679Z","response":" sky","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:41.964928002Z","response":" appears","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.080376927Z","response":" blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.197082451Z","response":" to","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.309908875Z","response":" us","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.419676198Z","response":" because","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.52530692Z","response":" of","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.624329141Z","response":" a","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.725207462Z","response":" phenomenon","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.826766183Z","response":" called","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:42.924925804Z","response":" Ray","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.028862825Z","response":"leigh","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.141743849Z","response":" scattering","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.320425187Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.426640409Z","response":" named","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.52803823Z","response":" after","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.629592551Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.728885672Z","response":" British","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.834318794Z","response":" physicist","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:43.935285215Z","response":" Lord","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.049440439Z","response":" Ray","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.155307862Z","response":"leigh","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.257485583Z","response":".","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.363339905Z","response":" He","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.462067623Z","response":" discovered","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.566671142Z","response":" that","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.672643762Z","response":" when","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.778044782Z","response":" sunlight","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.881093601Z","response":" enters","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:44.98386052Z","response":" Earth","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.086293039Z","response":"'s","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.187877657Z","response":" atmosphere","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.292611077Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.398258796Z","response":" it","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.503518516Z","response":" encounters","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.611737436Z","response":" tiny","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.713280455Z","response":" molecules","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.814858773Z","response":" of","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:45.920347893Z","response":" gases","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.020814011Z","response":" such","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.125800831Z","response":" as","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.22826835Z","response":" nitrogen","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.328753468Z","response":" and","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.433188188Z","response":" oxygen","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.535272007Z","response":".\n\n","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.637729326Z","response":"These","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.740624245Z","response":" gas","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.840901463Z","response":" molecules","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:46.946229783Z","response":" scatter","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.047443301Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.158124022Z","response":" light","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.265277642Z","response":" in","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.370128061Z","response":" all","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.475964281Z","response":" directions","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.5798255Z","response":",","done":false}{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.680482218Z","response":" but","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.781373137Z","response":" they","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.879549155Z","response":" scatter","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:47.993033476Z","response":" shorter","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.102438096Z","response":" (","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.203291515Z","response":"blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.308105334Z","response":")","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.411696454Z","response":" wavelengths","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.513615472Z","response":" more","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.615204591Z","response":" than","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.71910511Z","response":" longer","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.82229163Z","response":" (","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:48.921761648Z","response":"red","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.021251266Z","response":")","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.121801985Z","response":" wavelengths","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.223397004Z","response":".","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.331209724Z","response":" This","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.442343844Z","response":" is","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.585019771Z","response":" because","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.705397393Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.825608615Z","response":" smaller","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:49.963025241Z","response":" molecules","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:50.078111862Z","response":" are","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:50.194428883Z","response":" more","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:50.323840907Z","response":" effective","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:50.60927336Z","response":" at","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:50.840538603Z","response":" scattering","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:51.044260041Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:51.175701565Z","response":" blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:51.359158599Z","response":" light","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:51.594034242Z","response":" than","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:51.820704784Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:52.281164569Z","response":" larger","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:52.542569119Z","response":" molecules","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:52.677184545Z","response":".\n\n","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:52.796594868Z","response":"As","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:52.920132792Z","response":" a","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.099977227Z","response":" result","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.351723076Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.535188612Z","response":" when","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.666682237Z","response":" sunlight","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.78163946Z","response":" enters","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:53.905239083Z","response":" our","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.058695813Z","response":" atmosphere","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.267567654Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.489025297Z","response":" it","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.622062823Z","response":" is","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.740012345Z","response":" scattered","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:54.874145371Z","response":" in","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.004813297Z","response":" all","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.12252842Z","response":" directions","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.248260144Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.444142982Z","response":" giving","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.649850122Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.811343753Z","response":" sky","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:55.931928377Z","response":" its","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.0520861Z","response":" blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.172120823Z","response":" appearance","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.297780048Z","response":".","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.482415683Z","response":" The","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.675059121Z","response":" exact","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.848339354Z","response":" shade","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:56.975439279Z","response":" of","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.118281307Z","response":" blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.355813653Z","response":" can","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.568451094Z","response":" vary","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.739751927Z","response":" depending","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.865305752Z","response":" on","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:57.988667376Z","response":" atmospheric","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.117779701Z","response":" conditions","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.26709473Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.490771873Z","response":" such","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.701382214Z","response":" as","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.83578154Z","response":" pollution","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:58.961275064Z","response":" levels","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:59.097782991Z","response":" and","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:59.24775932Z","response":" temperature","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:59.439649457Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:59.647560697Z","response":" but","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T12:59:59.890917745Z","response":" overall","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:00.367369137Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:01.016379957Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:01.313481812Z","response":" sky","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:01.700973983Z","response":" remains","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:01.878696116Z","response":" blue","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.042520346Z","response":".\n\n","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.203188676Z","response":"It","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.38583991Z","response":"'s","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.537631138Z","response":" worth","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.687422865Z","response":" noting","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.835729993Z","response":" that","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:02.974849118Z","response":" during","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.121948545Z","response":" certain","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.248161469Z","response":" times","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.372707892Z","response":" of","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.487120313Z","response":" day","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.615138336Z","response":" or","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.732753558Z","response":" at","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.864681982Z","response":" high","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:03.984105204Z","response":" alt","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.106823927Z","response":"itudes","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.220818648Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.334663069Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.460696892Z","response":" sky","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.580111514Z","response":" can","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.704296037Z","response":" appear","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.819842059Z","response":" more","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:04.945050382Z","response":" red","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.069958205Z","response":" or","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.196488928Z","response":" even","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.406039867Z","response":" purple","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.522011088Z","response":" due","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.64096331Z","response":" to","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.765564833Z","response":" the","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:05.884748055Z","response":" scattering","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.013101579Z","response":" of","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.1289723Z","response":" light","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.250474422Z","response":" in","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.369119144Z","response":" different","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.491995667Z","response":" ways","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.608972888Z","response":".","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.733432411Z","response":" However","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.851776233Z","response":",","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:06.967879055Z","response":" these","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.086132676Z","response":" scenarios","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.217575101Z","response":" are","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.344529324Z","response":" relatively","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.466911447Z","response":" rare","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.665886383Z","response":" and","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:07.827181313Z","response":" usually","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.015870448Z","response":" occur","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.141287171Z","response":" under","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.254635792Z","response":" specific","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.379744715Z","response":" circumstances","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.489780436Z","response":".","done":false}
{"model":"llama3.2:1b","created_at":"2025-02-22T13:00:08.61143366Z","response":"","done":true,"done_reason":"stop","context":[128006,9125,128007,271,38766,1303,33025,2696,25,6790,220,2366,18,271,128009,128006,882,128007,271,10445,374,279,13180,6437,30,128009,128006,78191,128007,271,791,13180,8111,6437,311,603,1606,315,264,25885,2663,13558,64069,72916,11,7086,1306,279,8013,83323,10425,13558,64069,13,1283,11352,430,994,40120,29933,9420,596,16975,11,433,35006,13987,35715,315,45612,1778,439,47503,323,24463,382,9673,6962,35715,45577,279,3177,304,682,18445,11,719,814,45577,24210,320,12481,8,93959,810,1109,5129,320,1171,8,93959,13,1115,374,1606,279,9333,35715,527,810,7524,520,72916,279,6437,3177,1109,279,8294,35715,382,2170,264,1121,11,994,40120,29933,1057,16975,11,433,374,38067,304,682,18445,11,7231,279,13180,1202,6437,11341,13,578,4839,28601,315,6437,649,13592,11911,389,45475,4787,11,1778,439,25793,5990,323,9499,11,719,8244,11,279,13180,8625,6437,382,2181,596,5922,27401,430,2391,3738,3115,315,1938,477,520,1579,4902,21237,11,279,13180,649,5101,810,2579,477,1524,25977,4245,311,279,72916,315,3177,304,2204,5627,13,4452,11,1521,26350,527,12309,9024,323,6118,12446,1234,3230,13463,13],"total_duration":30775260423,"load_duration":3092495248,"prompt_eval_count":31,"prompt_eval_duration":800000000,"eval_count":192,"eval_duration":26881000000}

# Technical Uncertainty

Q Does bridge mode mean we can only accses Ollama API with another model in the docker compose?

A No, the host machine will be able to access it

Q: Which port is being mapped 8008 should be kept the same 

In this case 8008 is the port that host machine will access. the other other in the guest port (the port of the service inside container)

Q: If we pass the LLM_MODEL_Id to the ollama server will it download the model when on start?

It does not appear so. The ollama CLI might be running multiple APIs so you need to call the /pull api before trying to generat text

Q: Will the model be downloaded in the container? does that mean the ml model will be deleted when the container stops running?

A: The model will download into the container, and vanish when the container stop running. You need to mount a local drive and there is probably more work to be done.

Q: For LLM service which can text-generation it suggets it will only work with TGI/vLLM and all you have to do is have it running. Does TGI and vLLM have a stardarized API or is there code to detect which one is running? Do we have to really use Xeon or Guadi processor?

vLLM, TGI (Text Generation Inference), and Ollama all offer APIs with OpenAI compatibility, so in theory they should be interchangable.