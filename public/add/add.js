let citySelect = document.getElementById("city")
let provinceSelect = document.getElementById("province")
let addForm = document.forms["addForm"]

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
    if (addForm["province"].value && addForm["city"].value && addForm["goldar"].value && addForm["rhesus"].value && addForm["name"].value && addForm["age"].value && addForm["gender"].value && addForm["desc"].value  && addForm["hospitalname"].value && addForm["captcha"].value){
        //checkCaptcha
        return true
    }
    alert("harap isi dengan benar")
    return false
}