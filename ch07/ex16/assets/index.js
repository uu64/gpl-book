function input(v) {
  document.getElementById("answer").textContent += v
}

function ce() {
  document.getElementById("expr").textContent = ""
  document.getElementById("answer").textContent = ""
}

function enter() {
  var expr = document.getElementById("answer").textContent
  fetch("http://localhost:8080/calc/", {
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      expr: expr
    }),
  }).then(async res => {
    if(!res.ok) {
      throw await res.json()
    }
    return res.json()
  }).then(res => {
    document.getElementById("expr").textContent = res.expr + "="
    document.getElementById("answer").textContent = res.answer
  }).catch(err => {
    console.error(err.message)
    document.getElementById("answer").textContent = "Error!"
  })
}