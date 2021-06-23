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
    let age = addForm["age"].value
    let province = addForm["province"].value
    let city = addForm["city"].value
    let goldar = addForm["goldar"].value
    let rhesus = addForm["rhesus"].value
    let name = addForm["name"].value
    let gender = addForm["gender"].value
    let desc = addForm["desc"].value
    let hospitalname = addForm["hospitalname"].value
    let captcha = addForm["captcha"].value
    if (validateAge(+age) && province && city && goldar && rhesus && name && age && gender && desc  && hospitalname && captcha){
        return true
    }
    alert("harap isi dengan benar")
    return false
}

validateAge = (age)=>{
    let pattern = new RegExp(/^[\d]{1,3}$/)
    return pattern.test(age)
}