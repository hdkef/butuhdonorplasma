let cityidSelect = document.getElementById("cityid")
let provinceidSelect = document.getElementById("provinceid")
let addContact2 = document.getElementById("addcontact2")
let addForm = document.forms["addForm"]
let cp = document.getElementById("cp")
let selectedprovinceid
let provincename = document.getElementById("provincename")
let cityname = document.getElementById("cityname")

addContact2.onclick = (event)=>{
    var label = document.createElement('label')
    label.innerText = "contact person 2"
    var cpname2 = document.createElement('input')
    cpname2.name = "cpname2"
    cpname2.placeholder = "Nama Contact Person"
    var relation2 = document.createElement('input')
    relation2.placeholder = "Hubungan dengan pasien"
    relation2.name = "cprelation2"
    var tel2 = document.createElement('textarea')
    tel2.name = "cptel2"
    tel2.placeholder = "List cara menghubungi (WA,Telegram,Line)"
    cp.append(label,cpname2,relation2,tel2)
    addContact2.style.display = "none"
}

provinceidSelect.onchange = (event)=>{
    getCity(provinceidSelect,selectedprovinceid,cityidSelect) //ada di public.js
}

validateForm = (e)=>{
    let age = addForm["age"].value
    let provinceid = addForm["provinceid"].value
    let cityid = addForm["cityid"].value
    let goldar = addForm["goldar"].value
    let rhesus = addForm["rhesus"].value
    let name = addForm["name"].value
    let desc = addForm["desc"].value
    let hospitalname = addForm["hospitalname"].value
    let captcha = addForm["captcha"].value
    provincename.value = provinceidSelect.options[provinceidSelect.selectedIndex].innerHTML
    cityname.value = cityidSelect.options[cityidSelect.selectedIndex].innerHTML
    if (validateAge(+age) && provinceid && cityid && goldar && rhesus && name && age && desc  && hospitalname && captcha){
        return true
    }
    alert("harap isi form dengan benar")
    return false
}

validateAge = (age)=>{
    let pattern = new RegExp(/^[\d]{1,3}$/)
    return pattern.test(age)
}