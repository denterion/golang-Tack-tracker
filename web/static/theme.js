// theme.js
const toggle = document.getElementById('theme-toggle');
const body = document.body;

// Загружаем сохранённую тему при загрузке любой страницы
const currentTheme = localStorage.getItem('theme') || 'light';
if (currentTheme === 'dark') {
  body.classList.add('dark');
  if (toggle) toggle.checked = true;
}

if (toggle) {
  toggle.addEventListener('change', () => {
    if (toggle.checked) {
      body.classList.add('dark');
      localStorage.setItem('theme', 'dark');
    } else {
      body.classList.remove('dark');
      localStorage.setItem('theme', 'light');
    }
  });
}

