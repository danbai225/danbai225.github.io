// 建站时间统计
function show_date_time() {
  if ($("#span_dt_dt").length > 0) {
    window.setTimeout("show_date_time()", 1000);
    BirthDay = new Date("10/10/2019 16:20:09");
    today = new Date();
    timeold = today.getTime() - BirthDay.getTime();
    sectimeold = timeold / 1000;
    secondsold = Math.floor(sectimeold);
    msPerDay = 24 * 60 * 60 * 1000;
    e_daysold = timeold / msPerDay;
    daysold = Math.floor(e_daysold);
    e_hrsold = (e_daysold - daysold) * 24;
    hrsold = Math.floor(e_hrsold);
    e_minsold = (e_hrsold - hrsold) * 60;
    minsold = Math.floor((e_hrsold - hrsold) * 60);
    seconds = Math.floor((e_minsold - minsold) * 60);
    span_dt_dt.innerHTML =
      daysold + "天" + hrsold + "小时" + minsold + "分" + seconds + "秒";
  }
}

// 主题切换功能
function initThemeToggle() {
  // 获取当前主题
  const currentTheme = localStorage.getItem('theme') || 'light';
  document.documentElement.setAttribute('data-theme', currentTheme);
  
  // 创建主题切换按钮
  const themeToggle = $(`
    <div class="theme-toggle" title="切换主题">
      <span style="font-size: 20px;">${currentTheme === 'dark' ? '☀️' : '🌙'}</span>
    </div>
  `);
  
  $('body').append(themeToggle);
  
  // 主题切换事件
  themeToggle.click(function() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
    
    document.documentElement.setAttribute('data-theme', newTheme);
    localStorage.setItem('theme', newTheme);
    
    // 更新按钮图标
    $(this).find('span').text(newTheme === 'dark' ? '☀️' : '🌙');
    
    // 添加切换动画
    $('body').addClass('theme-transition');
    setTimeout(() => {
      $('body').removeClass('theme-transition');
    }, 300);
  });
}

// 返回顶部功能
function initBackToTop() {
  const backToTop = $(`
    <button class="back-to-top" title="返回顶部">
      <span style="font-size: 20px;">↑</span>
    </button>
  `);
  
  $('body').append(backToTop);
  
  // 滚动监听
  $(window).scroll(function() {
    if ($(this).scrollTop() > 300) {
      backToTop.addClass('visible');
    } else {
      backToTop.removeClass('visible');
    }
  });
  
  // 点击返回顶部
  backToTop.click(function() {
    $('html, body').animate({
      scrollTop: 0
    }, 600);
  });
}

// 页面加载动画
function initPageAnimation() {
  // 为所有卡片添加进入动画
  $('.item').each(function(index) {
    $(this).css({
      'opacity': '0',
      'transform': 'translateY(30px)'
    });
    
    setTimeout(() => {
      $(this).css({
        'opacity': '1',
        'transform': 'translateY(0)',
        'transition': 'all 0.6s cubic-bezier(0.4, 0, 0.2, 1)'
      });
    }, index * 100);
  });
}

// 搜索框增强
function enhanceSearch() {
  const searchBox = $("#searchBox");
  
  // 添加搜索图标
  searchBox.attr('placeholder', '🔍 搜索文章...');
  
  // 搜索建议功能（可以根据需要扩展）
  searchBox.on('input', function() {
    const query = $(this).val();
    if (query.length > 2) {
      // 这里可以添加搜索建议逻辑
      console.log('搜索:', query);
    }
  });
  
  // 搜索提交
  searchBox.keyup(function (e) {
    if (e.keyCode === 13) {
      const query = $(this).val();
      if (query.trim()) {
        window.location = "https://www.google.com/search?q=site%3Adanbai225.github.io+" + encodeURIComponent(query);
      }
    }
  });
}

// 懒加载增强
function initLazyLoading() {
  // 为图片添加懒加载类
  $('img').each(function() {
    if (!$(this).hasClass('lazyload')) {
      $(this).addClass('lazyload');
    }
  });
}

// 文章卡片交互增强
function enhanceArticleCards() {
  $('.item').hover(
    function() {
      // 鼠标进入时的动画
      $(this).find('.item-heading h4 a').css('color', 'var(--primary-color)');
    },
    function() {
      // 鼠标离开时的动画
      $(this).find('.item-heading h4 a').css('color', 'var(--text-color)');
    }
  );
  
  // 点击卡片跳转
  $('.item').click(function(e) {
    if (e.target.tagName !== 'A') {
      const link = $(this).find('.item-heading h4 a').attr('href');
      if (link) {
        window.location.href = link;
      }
    }
  });
}

// 主题切换过渡样式
function addThemeTransition() {
  const style = $(`
    <style>
      .theme-transition * {
        transition: background-color 0.3s ease, color 0.3s ease, border-color 0.3s ease !important;
      }
    </style>
  `);
  $('head').append(style);
}

// 页面性能优化
function optimizePerformance() {
  // 预加载关键资源
  const preloadLinks = [
    '/res/fonts/zpix.woff2'
  ];
  
  preloadLinks.forEach(href => {
    const link = document.createElement('link');
    link.rel = 'preload';
    link.href = href;
    link.as = href.includes('.woff') ? 'font' : 'style';
    if (href.includes('.woff')) {
      link.crossOrigin = 'anonymous';
    }
    document.head.appendChild(link);
  });
}

// 错误处理
function initErrorHandling() {
  window.addEventListener('error', function(e) {
    console.error('页面错误:', e.error);
  });
  
  // 图片加载失败处理
  $('img').on('error', function() {
    $(this).css({
      'background': 'linear-gradient(45deg, #f0f0f0, #e0e0e0)',
      'display': 'flex',
      'align-items': 'center',
      'justify-content': 'center'
    }).text('图片加载失败');
  });
}

// 主初始化函数
$(function () {
  // 基础功能
  show_date_time();
  
  // 现代化功能
  addThemeTransition();
  initThemeToggle();
  initBackToTop();
  enhanceSearch();
  enhanceArticleCards();
  initLazyLoading();
  optimizePerformance();
  initErrorHandling();
  
  // 页面加载完成后的动画
  setTimeout(() => {
    initPageAnimation();
  }, 100);
  
  // 添加页面加载完成的淡入效果
  $('body').css('opacity', '0').animate({
    opacity: 1
  }, 800);
});

// 页面离开时的动画
$(window).on('beforeunload', function() {
  $('body').fadeOut(200);
});
