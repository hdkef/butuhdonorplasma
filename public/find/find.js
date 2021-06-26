let cityidSelect = document.getElementById("cityid")
let provinceidSelect = document.getElementById("provinceid")
let findForm = document.forms["findForm"]
let selectedprovinceid

provinceidSelect.onchange = (event)=>{
    getCity(provinceidSelect,selectedprovinceid,cityidSelect) //ada di public.js
}

validateForm = ()=>{
    let provinceid = findForm["provinceid"].value
    let cityid = findForm["cityid"].value
    if (provinceid && cityid && cityid){
        return true
    }
    alert("harap isi dengan benar")
    return false
}