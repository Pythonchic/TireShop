document.addEventListener("DOMContentLoaded", function () {
    // Получаем все необходимые элементы
    const elements = {
        lowerSlider: document.getElementById("lower"),
        upperSlider: document.getElementById("upper"),
        inputMin: document.getElementById("input-min"),
        inputMax: document.getElementById("input-max"),
        track: document.getElementById("track"),
        seasonFilter: document.getElementById("season-filter"),
        widthFilter: document.getElementById("width-filter"),
        profileFilter: document.getElementById("profile-filter"),
        diametrFilter: document.getElementById("diametr-filter"),
        applyFilterBtn: document.getElementById("apply-filter"),
        productRows: document.querySelectorAll("tbody tr")
    };

    // Обновление позиции ползунков и трека
    function updateSlider() {
        let lowerVal = parseInt(elements.lowerSlider.value);
        let upperVal = parseInt(elements.upperSlider.value);

        // Защита от пересечения значений
        if (lowerVal > upperVal) {
            [lowerVal, upperVal] = [upperVal, lowerVal];
            elements.lowerSlider.value = lowerVal;
            elements.upperSlider.value = upperVal;
        }

        // Расчет позиции трека
        const min = parseInt(elements.lowerSlider.min);
        const max = parseInt(elements.lowerSlider.max);
        const leftPercent = ((lowerVal - min) / (max - min)) * 100;
        const rightPercent = 100 - ((upperVal - min) / (max - min)) * 100;

        elements.track.style.left = leftPercent + "%";
        elements.track.style.right = rightPercent + "%";

        // Обновление полей ввода
        elements.inputMin.value = lowerVal;
        elements.inputMax.value = upperVal;
    }

    // Обработка изменения значений в полях ввода
    function handleInputChange(e) {
        const input = e.target;
        let value = input.value === "" ? "" : parseInt(input.value);
        const minvalue = parseInt(elements.lowerSlider.min);
        const maxvalue = parseInt(elements.upperSlider.max);

        // Если поле пустое - не обновляем ползунок
        if (value === "") return;
        if (parseInt(value) < minvalue) return;

        if (parseInt(value) > maxvalue) {
            input.value = maxvalue;
            value = maxvalue;
        }

        // Обновление соответствующих ползунков
        if (input.id === "input-min") {
            elements.lowerSlider.value = value;
            if (value > parseInt(elements.upperSlider.value)) {
                elements.upperSlider.value = value;
            }
        } else {
            elements.upperSlider.value = value;
            if (value < parseInt(elements.lowerSlider.value)) {
                elements.lowerSlider.value = value;
            }
        }

        updateSlider();
    }

    // Обработка потери фокуса (если поле осталось пустым)
    function handleInputBlur(e) {
        const input = e.target;
        if (input.value === "") {
            const defaultValue = input.id === "input-min" ? input.min : input.max;
            input.value = defaultValue;

            if (input.id === "input-min") {
                elements.lowerSlider.value = defaultValue;
            } else {
                elements.upperSlider.value = defaultValue;
            }

            updateSlider();
        }
    }

    // Фильтрация товаров
    function applyFilters() {
        const selectedSeason = elements.seasonFilter.value.toLowerCase();
        const selectedWidth = elements.widthFilter.value.toLowerCase();
        const selectedProfile = elements.profileFilter.value.toLowerCase();
        const selectedDiametr = elements.diametrFilter.value.toLowerCase();

        let max = parseInt(elements.inputMax.value);
        let min = parseInt(elements.inputMin.value);
        const minvalue = parseInt(elements.lowerSlider.min);
        const maxvalue = parseInt(elements.lowerSlider.max);

        if (min < minvalue) {
            min = minvalue;
            elements.inputMin.value = minvalue;
        }

        if (max < minvalue) {
            max = maxvalue;
            elements.inputMax.value = maxvalue;
        }

        elements.productRows.forEach(function (row) {
            const cells = row.cells;
            const cost = parseInt(cells[1].textContent);
            const season = cells[6].textContent.toLowerCase();
            const width = cells[3].textContent.toLowerCase();
            const profile = cells[4].textContent.toLowerCase();
            const diametr = cells[2].textContent.toLowerCase();

            const priceMatch = cost <= max && cost >= min;
            const seasonMatch = selectedSeason === "all" || season === selectedSeason;
            const widthMatch = selectedWidth === "all" || width === selectedWidth;
            const profileMatch = selectedProfile === "all" || profile === selectedProfile;
            const diametrMatch = selectedDiametr === "all" || diametr === selectedDiametr;

            row.style.display = priceMatch && seasonMatch && widthMatch && profileMatch && diametrMatch ? "" : "none";
        });
    }

    // Назначаем обработчики событий
    elements.lowerSlider.addEventListener("input", updateSlider);
    elements.upperSlider.addEventListener("input", updateSlider);

    elements.inputMin.addEventListener("input", handleInputChange);
    elements.inputMax.addEventListener("input", handleInputChange);

    elements.inputMin.addEventListener("blur", handleInputBlur);
    elements.inputMax.addEventListener("blur", handleInputBlur);

    elements.applyFilterBtn.addEventListener("click", applyFilters);

    // Инициализация
    updateSlider();
    applyFilters();
});
