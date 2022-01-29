const elements = document.querySelectorAll(".icon");

function letterUsed(element) {
    element.classList.add('clicked');
    element.classList.remove('letter');
    console.log(element.querySelector("span").innerText);
}

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

document.addEventListener('keydown', event => {
    const key = event.key.toLowerCase();
    if (key.length === 1 && key >= 'a' && key <= 'z') {
        const element = [...document.querySelectorAll("span")]
            .find(element => element.innerText === key.toUpperCase()).parentNode;

        if (!element) return;
        letterUsed(element)
    }
});
