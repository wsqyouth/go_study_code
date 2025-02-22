document.addEventListener('DOMContentLoaded', function() {
  const uuidInput = document.getElementById('uuidInput');
  const convertButton = document.getElementById('convertButton');
  const errorMessage = document.getElementById('errorMessage');

  function isValidUUID(str) {
    // 检查带横线的UUID格式
    const uuidWithDashPattern = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i;
    // 检查不带横线的UUID格式
    const uuidWithoutDashPattern = /^[0-9a-f]{32}$/i;
    
    return uuidWithDashPattern.test(str) || uuidWithoutDashPattern.test(str);
  }

  function convertUUID(uuid) {
    if (uuid.includes('-')) {
      // 如果包含横线，则移除所有横线
      return uuid.replace(/-/g, '');
    } else {
      // 如果不包含横线，则添加横线
      return uuid.replace(/([0-9a-f]{8})([0-9a-f]{4})([0-9a-f]{4})([0-9a-f]{4})([0-9a-f]{12})/i, '$1-$2-$3-$4-$5');
    }
  }

  convertButton.addEventListener('click', function() {
    const inputValue = uuidInput.value.trim();
    
    if (!isValidUUID(inputValue)) {
      errorMessage.style.display = 'block';
      return;
    }
    
    errorMessage.style.display = 'none';
    const convertedValue = convertUUID(inputValue);
    uuidInput.value = convertedValue;

    // 自动选中转换后的文本，方便用户复制
    uuidInput.select();
  });

  // 当输入框内容变化时，隐藏错误信息
  uuidInput.addEventListener('input', function() {
    errorMessage.style.display = 'none';
  });
});