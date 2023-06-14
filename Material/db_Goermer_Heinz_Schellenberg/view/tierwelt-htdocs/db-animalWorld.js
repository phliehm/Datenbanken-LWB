function jsTest() {
    alert("kuckuck");
}

function jsPruef(name) {
    ausgabe = '';
    for (i = 0; i < 6; i++) {
        ausgabe = ausgabe + '\n' + (document.getElementById(name + '_' + i).value);
    }
    alert(ausgabe);
}

function Pruefen(id,table) {
    if (window.XMLHttpRequest) {
            // code for IE7+, Firefox, Chrome, Opera, Safari
            xmlhttp = new XMLHttpRequest();
        } else {
            // code for IE6, IE5
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            document.getElementById('jsAusgabe').innerHTML = this.responseText;
        }
    }
    val = document.getElementsByName(id)[0].value;
    xmlhttp.open("GET","jsQuery_valid.php?id=" + id + "&val=" + val + "&table=" + table, true);
    xmlhttp.send();
    //document.getElementById(elementName).style.display = 'block';
}
