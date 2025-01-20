
tailwind.config = {
    darkMode: false,
};

const body = document.body;

function createMenuButton() {
    const menuButton = document.createElement('button');
    menuButton.id = 'menu-toggle';
    menuButton.className = 'absolute top-4 left-4 menu-btn';
    menuButton.innerHTML = '<span class="material-icons">🏡</span>';

    menuButton.addEventListener('click', () => {
        const appDrawer = document.getElementById('app-drawer');
        appDrawer.classList.toggle('-translate-x-full');
    });

    return menuButton;
}

function createThemeButton() {
    const themeButton = document.createElement('button');
    themeButton.id = 'theme-toggle';
    themeButton.className = 'absolute top-4 right-4 btn theme-btn';
    themeButton.innerHTML = '<span id="theme-icon" class="theme-icon sun">☀️</span>';

    themeButton.addEventListener('click', () => {
        const isDark = body.classList.contains('theme-dark');
        body.classList.toggle('theme-light', isDark);
        body.classList.toggle('theme-dark', !isDark);

        const themeIcon = document.getElementById('theme-icon');
        themeIcon.classList.toggle('sun', !isDark);
        themeIcon.classList.toggle('moon', isDark);
        themeIcon.textContent = isDark ? '☀️' : '🌙';
    });

    return themeButton;
}

function initializeHeaderButtons() {
    const header = document.getElementById('header-buttons');
    header.appendChild(createMenuButton());
    header.appendChild(createThemeButton());
}

function enableCloseMenuButton() {
    const closeMenuButton = document.getElementById('close-menu');
    closeMenuButton.addEventListener('click', () => {
        const appDrawer = document.getElementById('app-drawer');
        appDrawer.classList.add('-translate-x-full');
    });
}

document.addEventListener('DOMContentLoaded', () => {
    initializeHeaderButtons();
    enableCloseMenuButton();
});

document.addEventListener('DOMContentLoaded', initializeHeaderButtons);