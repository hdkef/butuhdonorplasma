let deleteForm = document.forms["deleteForm"]


validateForm = (event)=>{
    let id = deleteForm["id"].value
    if (id){
        if (!sanitize(id)){
            return false //if input contain 'script'
        }
        return true
    }
    alert("id masih kosong")
    return false
}