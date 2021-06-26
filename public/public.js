getCity = (provinceidSelect,selectedprovinceid,cityidSelect)=>{
    let id = provinceidSelect.value
    if (selectedprovinceid != id) {
         //Jika provinsi yang dipilih adalah bukan provinsi semua, maka fetch city
        deleteChildren(cityidSelect) //hapus dulu atau reset city options
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
                cityidSelect.append(added);
            }
        })
    }
}

deleteChildren = (parent)=>{
    while (parent.firstChild) {
        console.log("remove")
        parent.removeChild(parent.firstChild);
    }
}