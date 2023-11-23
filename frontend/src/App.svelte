<script lang="ts">
  import { Time } from '../wailsjs/go/main/Timer.js'
  import { Quit } from '../wailsjs/runtime'
  import { onDestroy } from 'svelte';

  let cooldown: boolean = false
  let elapsed = 0;
	let duration = 1200000;

	let last_time = window.performance.now();
	let frame;

	(function update() {
		frame = requestAnimationFrame(update);

		const time = window.performance.now();
		elapsed += Math.min(time - last_time, duration - elapsed);

		last_time = time;
	})();

	onDestroy(() => {
		cancelAnimationFrame(frame);
	});

  $: if (elapsed - duration === 0) {
    Time()
    cooldown = true
  }

  function quitApp() {
    Quit()
  }

  function mConv(millis: number) {
    const minutes = Math.floor(millis / 60000);
    const seconds = ((millis % 60000) / 1000).toFixed(0);
    return minutes + ":" + (Number(seconds) < 10 ? '0' : '') + seconds;
  }
</script>

<main>
  {#if !cooldown}
    <div class="result" id="result">Timer</div>

    <label>
	    Elapsed time:
	    <progress value={elapsed / duration} />
    </label>

    <div>{mConv(elapsed)}</div>
  {:else}
    <div class="result" id="result">
      Timer complete. Look at something 20 feet away for 20s.
      <button class="button-4" role="button" on:click={quitApp}>Quit</button>
    </div>
  {/if}
</main>

<style>
  .button-4 {
    padding-top: 10px;
    appearance: none;
    background-color: #FAFBFC;
    border: 1px solid rgba(27, 31, 35, 0.15);
    border-radius: 6px;
    box-shadow: rgba(27, 31, 35, 0.04) 0 1px 0, rgba(255, 255, 255, 0.25) 0 1px 0 inset;
    box-sizing: border-box;
    color: #24292E;
    cursor: pointer;
    display: inline-block;
    font-family: -apple-system, system-ui, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
    font-size: 14px;
    font-weight: 500;
    line-height: 20px;
    list-style: none;
    padding: 6px 16px;
    position: relative;
    transition: background-color 0.2s cubic-bezier(0.3, 0, 0.5, 1);
    user-select: none;
    -webkit-user-select: none;
    touch-action: manipulation;
    vertical-align: middle;
    white-space: nowrap;
    word-wrap: break-word;
  }

  .button-4:hover {
    background-color: #F3F4F6;
    text-decoration: none;
    transition-duration: 0.1s;
  }

  .button-4:active {
    background-color: #EDEFF2;
    box-shadow: rgba(225, 228, 232, 0.2) 0 1px 0 inset;
    transition: none 0s;
  }

  .button-4:focus {
    outline: 1px transparent;
  }

  .button-4:before {
    display: none;
  }

  .button-4:-webkit-details-marker {
    display: none;
  }
  
  .result {
    line-height: 20px;
    font-size: large;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-top: -50px;
    margin-left: -50px;
    width: 100px;
    height: 100px;
  }

</style>
