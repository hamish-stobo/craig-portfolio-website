document.addEventListener('DOMContentLoaded', () => {
    // Select all anchor links
    const links = document.querySelectorAll('a[href^="#"]');
    
    links.forEach(link => {
      link.addEventListener('click', (e) => {
        e.preventDefault();
        
        const targetId = link.getAttribute('href').substring(1);
        const targetElement = document.getElementById(targetId);
  
        if (targetElement) {
          window.scrollTo({
            top: targetElement.offsetTop,
            behavior: 'smooth'
          });
        }
      });
    });
  });

  const intervalTime = 2000;
  const wrapper = document.querySelector('.carousel-wrapper');
  const items = document.querySelectorAll('.carousel-item');
  const itemCount = items.length;
  let index = 0;

  function moveToNextSlide() {
      index = (index + 1) % itemCount;
      const offset = -index * 100;
      wrapper.style.transform = `translateX(${offset}%)`;
  }

    setInterval(moveToNextSlide, intervalTime);