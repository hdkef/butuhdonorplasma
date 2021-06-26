let modal = document.getElementById("modal")

openModal = (cpname1,cptel1,cprelation1, cpname2, cptel2, cprelation2)=>{
    var cpname1 = document.createElement('h1');
    cpname1.innerHTML = cpname1
}

reset = (modal)=>{
    while (modal.firstChild) {
        modal.removeChild(modal.firstChild);
    }
}