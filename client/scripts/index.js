let elements = document.querySelectorAll(".icon")

function letterUsed(element) {
    const letter = element.querySelector("span").innerText;
    if (element.classList.contains('clicked')) {
        console.log('already clicked ' + letter)
    } else {
        const urlParams = new URLSearchParams(window.location.search);
        urlParams.set('word_text', letter);
        window.location.search = urlParams;
        element.classList.add('clicked')
        element.classList.remove('letter')
        element.querySelector('.tooltip').style.display = "none";
        console.log(letter)
    }
}

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
        element1.classList.remove('border_letter')
        letterUsed(element1)
    }
});

document.addEventListener('keydown', event => {
    const key = event.key.toLowerCase();
    if (document.targetElement.tagName !== "form") {
        if (key.length === 1 && key >= 'a' && key <= 'z') {
            const element = [...document.querySelectorAll("span")]
                .find(element => element.innerText === key.toUpperCase()).parentNode;

            if (!element) return;
            if (document.hasFocus()) {
                if (!element.classList.contains('clicked')) {
                    element.classList.add('border_letter')
                }
            }
        }
    }
});

elements.forEach(element => {
    element.addEventListener("click", () => letterUsed(element), {once: true});
    element.addEventListener('mouseenter', () => {
        if (element.classList.contains('clicked')) return;
        const tooltip = document.querySelector('.tooltip');
        tooltip.style.visibility = 'visible';
        tooltip.remove();
        element.appendChild(tooltip);
        tooltip.classList.add('smooth-spawn');
    });

    element.addEventListener('mouseleave', () => {
        const tooltip = element.querySelector('.tooltip');
        tooltip.style.visibility = 'hidden';
        tooltip.classList.remove('smooth-spawn');
    });
});
