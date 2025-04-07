# gin-easypage

# Pagination Utils per Gin + GORM

Questo pacchetto contiene due utility per la gestione della **paginazione server-side** in applicazioni web costruite con [Gin](https://github.com/gin-gonic/gin) e [GORM](https://gorm.io/). Entrambe le funzioni permettono di restituire dataset paginati e includono nell'header HTTP il totale degli elementi presenti, utile per la gestione lato client.

---

## Funzionalità principali

- Aggiunta automatica di `LIMIT` e `OFFSET` in base ai parametri `page` e `page_size`.
- Calcolo del numero totale di righe e inserimento dell’header `X-Total-Count` nella risposta HTTP.
- Supporto a condizioni personalizzate (`WHERE`, `JOIN`) e ricerca globale (`globalSearch`).
- Compatibilità sia con query GORM che con query SQL personalizzate.

---

## Parametri URL supportati

- `page`: numero della pagina da restituire (0-based, cioè la prima pagina è 0).
- `page_size`: numero di elementi per pagina.

Questi parametri vanno passati nella query string della richiesta HTTP. Ad esempio:

GET /api/users?page=1&page_size=20

---

## Funzione `Paginate`

Questa funzione è pensata per essere usata con query GORM standard. Permette di specificare:
- La tabella di riferimento.
- Eventuali condizioni (`WHERE`) da applicare.
- Join SQL personalizzati.
- Una stringa per la ricerca globale.

### Quando usarla

Usa questa funzione quando stai eseguendo una query GORM classica su una singola tabella o con join dichiarati, e vuoi paginare il risultato.

### Header HTTP generato

Oltre al risultato paginato, la risposta conterrà un header:

X-Total-Count: <numero_totale_risultati>

Questo header può essere usato dal frontend per sapere quante pagine sono disponibili.

---

## Funzione `PaginateCustomQuery`

Questa funzione è pensata per essere usata quando si lavora con **query SQL personalizzate** (raw SQL). Restituisce una stringa contenente `LIMIT` e `OFFSET`, calcolati in base ai parametri `page` e `page_size` ricevuti nella query string.

### Quando usarla

Usa questa funzione quando:
- Hai una query SQL complessa che non puoi rappresentare facilmente con GORM.
- Devi unire più tabelle o sottoquery.
- Vuoi gestire direttamente la query raw, ma avere comunque la paginazione automatica e l’header `X-Total-Count`.

### Output della funzione

Restituisce una stringa che può essere concatenata alla tua query SQL per aggiungere paginazione, come ad esempio:

LIMIT 20 OFFSET 40

Anche in questo caso, l’header `X-Total-Count` viene automaticamente inserito nella risposta HTTP.

---

## Comportamento con `page_size` non definito

Se il parametro `page_size` non è presente o vale `0`, la funzione **non applica nessuna limitazione** al numero di risultati. In questo caso:
- Nessun `LIMIT/OFFSET` sarà applicato.
- L’intero dataset sarà restituito.

---

## Requisiti

- Gin framework
- GORM ORM

---

## Note finali

Le funzioni sono progettate per essere facilmente integrabili nei controller del tuo progetto. Per una perfetta integrazione lato frontend, si dovranno utilizzare i valori `X-Total-Count`, `page` e `page_size` per costruire interfacce paginabili dinamicamente.

---

## Licenza

Questo pacchetto è rilasciato sotto licenza MIT.