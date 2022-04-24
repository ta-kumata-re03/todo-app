<template>
  <div class="create-todo">
    <div class="create-todo__card">
      <form class="form">
        <h1>ToDoリストを作成する</h1>

        <div class="form__input-container">
          <div class="form__input-item">
            <label for="title">タイトル</label>
            <input
              type="text"
              id="title"
              name="title"
              maxlength="30"
              placeholder="買い物をする"
              required
            />
          </div>
          <div class="form__input-item">
            <label for="expire_date">有効期限</label>
            <input
              type="date"
              id="expire_date"
              name="expire_date"
              :min="expireStartDate"
              required
            />
          </div>
          <div class="form__input-item">
            <label for="detail">詳細</label>
            <textarea id="detail" name="detail" placeholder="人参を買う" />
          </div>
        </div>

        <button class="form__create-button">登録する</button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
const expireStartDate = computed(() => {
  const now = new Date();
  const year = now.getFullYear();
  const month = (now.getMonth() + 1).toString().padStart(2, "0");
  const date = now.getDate().toString().padStart(2, "0");
  return `${year}-${month}-${date}`;
});
</script>

<style scoped lang="scss">
.create-todo {
  width: 100%;
  height: 100vh;
  display: grid;
  place-items: center;
  background-color: #efeeea;

  &__card {
    height: 60%;
    padding: 36px;
    border-radius: 20px;
    background-color: #ffffff;
    box-shadow: 2px 2px 6px rgba(0, 0, 0, 0.4);
  }
}

.form {
  display: flex;
  flex-direction: column;
  gap: 8px;

  &__input-container {
    display: flex;
    flex-direction: column;
  }

  &__input-item {
    display: flex;
    flex-direction: column;
  }

  &__create-button {
    padding: 10px 20px;
    background-color: #1c273e;
    color: #efeeea;
    border: none;
  }
}
</style>
