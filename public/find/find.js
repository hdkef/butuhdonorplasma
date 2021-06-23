let citySelect = document.getElementById("city")
let provinceSelect = document.getElementById("province")
let goldarSelect = document.getElementById("goldar")
let rhesusSelect = document.getElementById("rhesus")
let findForm = document.forms["findForm"]

provinceSelect.onchange = (event)=>{
    let id = provinceSelect.value
    fetch('/getcity',{
        method:'POST',
        body:JSON.stringify({id:id})
    }).then((res)=>{
        return res.json()
    }).then((cities)=>{
        for (let city of cities) {
            var added = document.createElement('option');
            added.value = city.id;
            added.innerHTML = city.name;
            citySelect.append(added);
        }
    })
}

validateForm = ()=>{
    if (findForm["province"].value && findForm["city"].value && findForm["goldar"].value && findForm["rhesus"].value){
        return true
    }
    alert("harap isi dengan benar")
    return false
}