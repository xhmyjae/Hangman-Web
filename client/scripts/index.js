const elements = document.querySelectorAll('.icon');
const inputs = [...document.querySelectorAll('p.player_input')];
const lettersEntered = inputs.map(letter => letter.innerText.replace('_', '')).join('');

function letterUsed(element) {
    const letter = element.querySelector('span').innerText;
    const urlParams = new URLSearchParams(window.location.search);
    urlParams.set('word_text', letter);
    window.location.search = urlParams.toString();
    element.classList.add('clicked');
    element.classList.remove('letter');
    element.querySelector('.tooltip').style.display = 'none';
}

elements.forEach(element => {
    element.addEventListener('click', () => letterUsed(element));
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
        if (!tooltip) return;
        tooltip.style.visibility = 'hidden';
        tooltip.classList.remove('smooth-spawn');
    });

    if (lettersEntered.includes(element.querySelector('span').innerText.toLowerCase())) {
        element.classList.add('clicked');
        element.classList.remove('letter');
        const tooltip = element.querySelector('.tooltip');
        if (!tooltip) return;
        tooltip.style.display = 'none';
        element.removeEventListener('click', () => letterUsed(element));
    }
});