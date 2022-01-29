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

document.addEventListener('keyup', function (event) {
    let key = event.key.toLowerCase();
    if (key >= 'a' && key <= 'z') {
        console.log(key)
        let element1 = [...document.querySelectorAll("span")].find(element => {
            return element.innerText === key.toUpperCase();
        }).parentNode;
        console.log(element1)
        if (!element1) {
            return
        }
        letter_used(element1)
        element1.classList.remove('border_letter')
    }
});

document.addEventListener('keydown', function (event) {
    let letter = document.querySelectorAll(".letter");
    let key = event.key.toLowerCase();
    if (key >= 'a' && key <= 'z') {
        let element1 = [...document.querySelectorAll("span")].find(element => {
            return element.innerText === key.toUpperCase();
        }).parentNode;
        if (!element1) {
            return
        }
        if (document.hasFocus()) {
            if (!element1.classList.contains('clicked')) {
                element1.classList.add('border_letter')
            }
        }
    }
});
