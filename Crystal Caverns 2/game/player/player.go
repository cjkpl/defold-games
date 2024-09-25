components {
  id: "player"
  component: "/game/player/player.script"
}
components {
  id: "dust"
  component: "/game/player/dust.particlefx"
  position {
    y: -8.0
  }
}
components {
  id: "bubbles"
  component: "/game/player/bubbles.particlefx"
}
components {
  id: "splash"
  component: "/game/player/splash.particlefx"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"static\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/core/game.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"world\"\n"
  "mask: \"boom\"\n"
  "mask: \"enemy\"\n"
  "mask: \"break\"\n"
  "mask: \"prox\"\n"
  "mask: \"belt\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: -3.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: 0.5\n"
  "      y: 2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 1\n"
  "    count: 1\n"
  "  }\n"
  "  data: 4.0\n"
  "  data: 3.5\n"
  "}\n"
  ""
}
embedded_components {
  id: "boom_factory"
  type: "factory"
  data: "prototype: \"/game/player/boom.go\"\n"
  ""
}
