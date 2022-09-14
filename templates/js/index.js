function httpGet(theUrl) {
  let xmlHttpReq = new XMLHttpRequest();
  xmlHttpReq.open("GET", theUrl, false); 
  xmlHttpReq.send(null);
  return xmlHttpReq.responseText;
}

const adviceBTN = document.getElementById("adviceBTN");
const secondb = document.getElementById("secondb");

adviceBTN.addEventListener("click", () => {
    var adviceObj = httpGet('http://advicehub.onrender.com/suradvice');
    var parsed = JSON.parse(adviceObj);

    var description = parsed.data.description;
    var author = parsed.data.author;

    console.log(author)

    secondb.innerHTML = `<p class="secondp">" ${description} "<br>- ${author}</p>`
    secondb.style.color = "white"
    secondb.style.fontFamily = "'Ms Madi', cursive";
    secondb.style.fontSize = "24px"
});