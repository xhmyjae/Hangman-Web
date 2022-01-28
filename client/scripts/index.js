let elements = document.querySelectorAll(".icon");

function letter_used(element) {
    if (element.classList.contains('clicked')) {
        console.log('already clicked ' + (element.querySelector("span").innerText))
    } else {
        element.classList.add('clicked')
        element.classList.remove('letter')
        element.querySelector('.tooltip').style.display = "none";
        console.log(element.querySelector("span").innerText)
    }
}

elements.forEach(element => {
    element.addEventListener("click", () => {
        letter_used(element);
    })
})

document.addEventListener('keydown', function (event) {
    if ((event.key >= 'a' && event.key <= 'z') || (event.key >= 'A' && event.key <= 'Z')) {
        console.log(event.key)
        // letter_used(event.key) //need to translate it (element = key)
    }
});
