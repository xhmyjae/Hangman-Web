document.addEventListener('keyup', event => {
    if (document.activeElement.tagName.toLowerCase() === 'input') return;
    const key = event.key.toLowerCase();
    if (key >= 'a' && key <= 'z') {
        const el = [...document.querySelectorAll('span')].find(element => element.innerText === key.toUpperCase()).parentNode;
        if (!el || el.parentElement.classList.contains('clicked')) return;
        el.classList.remove('border_letter');
        letterUsed(el);
    }
});

document.addEventListener('keydown', event => {
    if (document.activeElement.tagName.toLowerCase() === 'input') return;
    const key = event.key.toLowerCase();
    if (key.length === 1 && key >= 'a' && key <= 'z') {
        const element = [...document.querySelectorAll('span')].find(element => element.innerText === key.toUpperCase()).parentNode;

        if (!element) return;
        if (document.hasFocus()) {
            if (!element.classList.contains('clicked')) {
                element.classList.add('border_letter');
            }
        }
    }
});

const input = document.querySelector('input.word_text');
input.addEventListener('beforeinput', (event) => {
    if (event.data < 'a' || event.data > 'z') event.preventDefault();
});
