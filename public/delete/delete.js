let deleteForm = document.forms["deleteForm"]


validateForm = (event)=>{
    let id = deleteForm["id"].value
    if (id){
        return true
    }
    alert("id masih kosong")
    return false
}