let elements = document.querySelectorAll(".icon");
elements.forEach(element => {
    element.addEventListener("click", () => {
        element.classList.add('clicked')
        element.classList.remove('letter')
        element.querySelector('.tooltip').style.display = "none";
        console.log(element.querySelector("span").innerText)
    })
})